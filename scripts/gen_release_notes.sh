#!/bin/bash

# generate_release_notes.sh - Generate release_notes.md for go-xbuild-go releases
# Usage: ./generate_release_notes.sh [release_message]

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if VERSION file exists
if [[ ! -f VERSION ]]; then
    echo -e "${RED}ERROR: VERSION file not found${NC}"
    exit 1
fi

# Read version from VERSION file
VERSION=$(cat VERSION | tr -d '\n' | tr -d '\r')
echo -e "${GREEN}Generating release notes for version: ${VERSION}${NC}"

# Extract project name from current directory or first binary
PROJECT_NAME=$(basename "$(pwd)")
BIN_DIR="bin"

if [[ ! -d "$BIN_DIR" ]]; then
    echo -e "${RED}ERROR: bin directory not found${NC}"
    exit 1
fi

# Try to detect project name from binaries
FIRST_BINARY=$(ls "$BIN_DIR"/*.tar.gz 2>/dev/null | head -1 | xargs basename)
if [[ -n "$FIRST_BINARY" ]]; then
    PROJECT_NAME=$(echo "$FIRST_BINARY" | sed -E "s/-$VERSION.*//")
fi

echo -e "${YELLOW}Project name: ${PROJECT_NAME}${NC}"

# Release message (can be passed as argument or default)
RELEASE_MSG="${1:-Bug fixes and improvements}"

# Get list of binaries
CHECKSUMS_FILE="${PROJECT_NAME}-${VERSION}-checksums.txt"

# Find a sample Windows and Linux binary for examples
WINDOWS_BINARY=$(ls "$BIN_DIR"/${PROJECT_NAME}-${VERSION}-windows-amd64.d.zip 2>/dev/null || echo "")
LINUX_BINARY=$(ls "$BIN_DIR"/${PROJECT_NAME}-${VERSION}-linux-amd64.d.tar.gz 2>/dev/null || echo "")

# Generate release_notes.md
OUTPUT_FILE="/tmp/release_notes.md"

cat > "$OUTPUT_FILE" << EOF
# Release ${VERSION}

${RELEASE_MSG}

Please look at [ChangeLog](ChangeLog.md) for what has changed in the current version.

Please do not forget to verify Checksums before installing.

\`\`\`bash
$ sha256sum -c ${CHECKSUMS_FILE}
EOF

# Add checksums verification output (simulated)
if [[ -f "$BIN_DIR/$CHECKSUMS_FILE" ]]; then
    while IFS= read -r line; do
        if [[ -n "$line" ]]; then
            FILENAME=$(echo "$line" | awk '{print $2}')
            echo "${FILENAME}: OK" >> "$OUTPUT_FILE"
        fi
    done < "$BIN_DIR/$CHECKSUMS_FILE"
fi

cat >> "$OUTPUT_FILE" << 'EOF'
```

**Note for Windows users:** Microsoft Edge may flag Go binaries as potentially harmful due to false positives in its virus detection software. This is a known issue with Go-compiled executables. Please use Chrome or Firefox to download, or build from source if you prefer.

Sample content of binary archives is shown below. The `tar` command is available in Windows 10 (1803) and later, or you can use the GUI (right-click → Extract All). After extracting, copy/rename the binary somewhere in your PATH.

EOF

# Add Windows binary listing example if available
if [[ -n "$WINDOWS_BINARY" ]]; then
    cat >> "$OUTPUT_FILE" << EOF
\`\`\`bash
➤ tar -tf $(basename "$WINDOWS_BINARY")
EOF
    
    # Generate example output using tar
    tar -tf "$WINDOWS_BINARY" 2>/dev/null >> "$OUTPUT_FILE" || {
        echo "${PROJECT_NAME}-${VERSION}-windows-amd64.d/LICENSE" >> "$OUTPUT_FILE"
        echo "${PROJECT_NAME}-${VERSION}-windows-amd64.d/README.md" >> "$OUTPUT_FILE"
        echo "${PROJECT_NAME}-${VERSION}-windows-amd64.d/platforms.txt" >> "$OUTPUT_FILE"
        echo "${PROJECT_NAME}-${VERSION}-windows-amd64.d/${PROJECT_NAME}-${VERSION}-windows-amd64.exe" >> "$OUTPUT_FILE"
        echo "${PROJECT_NAME}-${VERSION}-windows-amd64.d/${PROJECT_NAME}.1" >> "$OUTPUT_FILE"
    }
    
    cat >> "$OUTPUT_FILE" << 'EOF'
```

EOF
fi

# Add Linux binary listing example if available
if [[ -n "$LINUX_BINARY" ]]; then
    cat >> "$OUTPUT_FILE" << EOF
\`\`\`bash
➤ tar -tvf $(basename "$LINUX_BINARY")
EOF
    
    # Generate example output
    tar -tvf "$LINUX_BINARY" 2>/dev/null >> "$OUTPUT_FILE" || {
        echo "-rw-r--r--  0 $(whoami) staff    1072 $(date '+%b %d %H:%M') ${PROJECT_NAME}-${VERSION}-linux-amd64.d/LICENSE" >> "$OUTPUT_FILE"
        echo "-rw-r--r--  0 $(whoami) staff   43936 $(date '+%b %d %H:%M') ${PROJECT_NAME}-${VERSION}-linux-amd64.d/README.md" >> "$OUTPUT_FILE"
        echo "-rw-r--r--  0 $(whoami) staff     903 $(date '+%b %d %H:%M') ${PROJECT_NAME}-${VERSION}-linux-amd64.d/platforms.txt" >> "$OUTPUT_FILE"
        echo "-rwxr-xr-x  0 $(whoami) staff 4763832 $(date '+%b %d %H:%M') ${PROJECT_NAME}-${VERSION}-linux-amd64.d/${PROJECT_NAME}-${VERSION}-linux-amd64" >> "$OUTPUT_FILE"
        echo "-rw-r--r--  0 $(whoami) staff   46650 $(date '+%b %d %H:%M') ${PROJECT_NAME}-${VERSION}-linux-amd64.d/${PROJECT_NAME}.1" >> "$OUTPUT_FILE"
    }
    
    cat >> "$OUTPUT_FILE" << 'EOF'
```

EOF
fi

# Add installation instructions
cat >> "$OUTPUT_FILE" << EOF
# Install

## Extract (unzip or untar)

Example:

\`\`\`bash
➤ tar -xvf ${PROJECT_NAME}-${VERSION}-linux-amd64.d.tar.gz
x ${PROJECT_NAME}-${VERSION}-linux-amd64.d/LICENSE
x ${PROJECT_NAME}-${VERSION}-linux-amd64.d/README.md
x ${PROJECT_NAME}-${VERSION}-linux-amd64.d/platforms.txt
x ${PROJECT_NAME}-${VERSION}-linux-amd64.d/${PROJECT_NAME}-${VERSION}-linux-amd64
x ${PROJECT_NAME}-${VERSION}-linux-amd64.d/${PROJECT_NAME}.1
\`\`\`

\`\`\`bash
➤ sudo /bin/cp ${PROJECT_NAME}-${VERSION}-linux-amd64.d/${PROJECT_NAME}-${VERSION}-linux-amd64 \\
     /usr/local/bin/${PROJECT_NAME}
➤ sudo cp ${PROJECT_NAME}-${VERSION}-linux-amd64.d/${PROJECT_NAME}.1 \\
        /usr/local/share/man/man1/
➤ ${PROJECT_NAME} -version
${PROJECT_NAME} ${VERSION}
https://github.com/muquit/${PROJECT_NAME}
➤ man ${PROJECT_NAME}
\`\`\`

Cross-compiled and released with [go-xbuild-go](https://github.com/muquit/go-xbuild-go)
EOF

echo -e "${GREEN}✓ Release notes generated: ${OUTPUT_FILE}${NC}"
echo -e "${YELLOW}Note: Review and edit the release message and any details as needed${NC}"
