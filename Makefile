README_ORIG=./docs/README.md
README=./README.md
BINARY=./ssl-tls-verify
VERSION := $(shell cat VERSION)
LDFLAGS := -ldflags "-w -s -X 'github.com/muquit/ssl-tls-verify/pkg/version.Version=$(VERSION)'"
BUILD_OPTIONS = -trimpath
MARKDOWN_TOC_PROG=markdown-toc-go
GLOSSARY_FILE=./docs/glossary.txt
MAIN_MD=./docs/main.md
MAN_PAGE=./docs/ssl-tls-verify.1


all: build build_all doc man_page

build:
	@echo "Compiling ..."
	go build $(BUILD_OPTIONS) $(LDFLAGS) -o $(BINARY)

# Linux/Mac
install: build doc man_page
	sudo cp $(BINARY) /usr/local/bin
	sudo cp $(MAN_PAGE) /usr/local/share/man/man1

build_all:
	@/bin/rm -rf ./bin
	@echo "*** Cross Compiling ..."
	go-xbuild-go -build-args '$(BUILD_OPTIONS) $(LDFLAGS)' \
		-additional-files $(MAN_PAGE)

doc: gen_files 
	@echo "*** Generating README.md with TOC ..."
	@touch $(README)
	@chmod 600 $(README)
	$(MARKDOWN_TOC_PROG) -i $(MAIN_MD) -o $(README) --glossary ${GLOSSARY_FILE} -f
	@chmod 444 $(README)
	$(MARKDOWN_TOC_PROG) -no-credit -i docs/ChangeLog.md -o ./ChangeLog.md --glossary ${GLOSSARY_FILE} -f

man_page:
	pandoc README.md -s -t man \
  -V title="ssl-tls-verify" \
  -V section="1" \
  -V header="User Commands" \
  -o $(MAN_PAGE)

test:
	go test ./pkg/utils/

gen_files: 
	@./scripts/mksf.sh

clean:
	@/bin/rm -rf ./bin
	@/bin/rm -f $(BINARY)
