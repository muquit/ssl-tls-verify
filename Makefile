README_ORIG=./docs/README.md
README=./README.md
BINARY=./ssl-tls-verify
VERSION := $(shell cat VERSION)
LDFLAGS := -ldflags "-w -s -X 'github.com/muquit/ssl-tls-verify/pkg/version.Version=$(VERSION)'"
BUILD_OPTIONS = -trimpath
MARKDOWN_TOC_PROG=markdown-toc-go
GLOSSARY_FILE=./docs/glossary.txt
SF=./docs/synopsis.txt
VF=./docs/version.md
BADGEF=./docs/badges.md
MAIN_MD=./docs/main.md
SAMPLE=./docs/sample_output.md


all: build build_all doc

build:
	@echo "Compiling ..."
	go build $(BUILD_OPTIONS) $(LDFLAGS) -o $(BINARY)

build_all:
	@/bin/rm -rf ./bin
	@echo "*** Cross Compiling ...."
	go-xbuild-go -build-args '$(BUILD_OPTIONS) $(LDFLAGS)'

doc: gen_files
	@echo "*** Generating README.md with TOC ..."
	@touch $(README)
	@chmod 600 $(README)
	$(MARKDOWN_TOC_PROG) -i $(MAIN_MD) -o $(README) --glossary ${GLOSSARY_FILE} -f
	@chmod 444 $(README)
	$(MARKDOWN_TOC_PROG) -no-credit -i docs/ChangeLog.md -o ./ChangeLog.md --glossary ${GLOSSARY_FILE} -f


gen_files: build gen_synopsis ver sample_output

gen_synopsis: 
	echo '## Synopsis' > $(SF)
	echo '```' >> $(SF)
	/bin/ls -lt $(BINARY)
	$(BINARY) -h >> $(SF) 2>&1
	echo '```' >> $(SF)

ver:
	echo "## Latest Version ($(VERSION))" > $(VF)
	echo "The current version is $(VERSION)" >> $(VF)
	echo "Please look at @CHANGELOG@ for what has changed in the current version.">> $(VF)

sample_output:
	MAILSEND_MASK_IP=1
	@echo "### Sample output" > ./docs/sample_output.md
	@echo "#### google.com HTTPS port 443" >> $(SAMPLE)
	@echo '```bash' >> $(SAMPLE)
	@echo "$(BINARY) --host google.com --port 443" >> $(SAMPLE)
	@echo '```' >> $(SAMPLE)
	@echo '```bash' >> $(SAMPLE)
	@MAILSEND_MASK_IP=1 $(BINARY) --host google.com --port 443 >> $(SAMPLE)
	@echo '```' >> $(SAMPLE)

	@echo "#### Server using self signed certificate" >> $(SAMPLE)
	@echo '```bash' >> $(SAMPLE)
	@echo '# Server using self-signed certificate' >> $(SAMPLE)
	@echo "$(BINARY) -host 127.0.0.1 -port 8881" >> $(SAMPLE)
	@echo '```' >> $(SAMPLE)
	@echo '```bash' >> $(SAMPLE)
	@MAILSEND_MASK_IP=1 $(BINARY) -host 127.0.0.1 -port 8881 >> $(SAMPLE) 2>&1
	@echo '```' >> $(SAMPLE)
	@echo "#### Skip Verification" >> $(SAMPLE)
	@echo '```bash' >> $(SAMPLE)
	@echo '# Skip verification' >> $(SAMPLE)
	@echo "$(BINARY) -host 127.0.0.1 -port 8881 -skip-verify" >> $(SAMPLE)
	@echo '```' >> $(SAMPLE)
	@echo '```bash' >> $(SAMPLE)
	@MAILSEND_MASK_IP=1 $(BINARY) -host 127.0.0.1 -port 8881 -skip-verify >> $(SAMPLE) 2>&1
	@echo '```' >> $(SAMPLE)
	@echo "#### SMTP with StartTLS" >> $(SAMPLE)
	@echo '```bash' >> $(SAMPLE)
	@echo "$(BINARY) --host smtp.gmail.com --port 587 --starttls" >> $(SAMPLE)
	@echo '```' >> $(SAMPLE)
	@echo '```bash' >> $(SAMPLE)
	@MAILSEND_MASK_IP=1 $(BINARY) --host smtp.gmail.com --port 587 --starttls >> $(SAMPLE) 2>&1
	@echo '```' >> $(SAMPLE)



clean:
	@/bin/rm -rf ./bin $
	@/bin/rm -f $(BINARY)
