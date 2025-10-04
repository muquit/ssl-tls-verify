## Installation

### Download pre-compiled binaries

Please download pre-compiled binaries from @RELEASES@ page. After extracting
the archive, copy/rename the binary to somewhere in your path.

### Build from source
```bash
go build -o ssl-tls-verify main.go
# or
make build
```

If you're on Linux/MacOS:
```bash
make install
```
It will compile and do the following:

```bash
    sudo cp ssl-tls-verify /usr/local/bin
    sudo cp docs/ssl-tls-verify.1 /usr/local/share/man/man1
```

### Install directly
```bash
go install github.com/muquit/ssl-tls-verify@latest
```
