README_ORIG=./docs/README.md
README=./README.md
BINARY=./ssl-tls-verify
VERSION := $(shell cat VERSION)
LDFLAGS := -ldflags "-w -s -X 'github.com/muquit/ssl-tls-verify/pkg/version.Version=$(VERSION)'"
BUILD_OPTIONS = -trimpath
MARKDOWN_TOC_PROG=markdown-toc-go
GLOSSARY_FILE=./docs/glossary.txt
MAIN_MD=./docs/main.md


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


gen_files: 
	@./scripts/mksf.sh

clean:
	@/bin/rm -rf ./bin
	@/bin/rm -f $(BINARY)
