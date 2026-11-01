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


.PHONY: all build install build_all release check_github_token doc docs man_page test gen_files clean \
	release_notes

all: build build_all doc man_page docs

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

# publish the GitHub release from ./bin using go-xbuild-go. It also
# pushes the homebrew formula. Requires GITHUB_TOKEN
release: check_github_token
	go-xbuild-go -release

release_notes:
	./scripts/gen_release_notes.sh
	/bin/cp -fv /tmp/release_notes.md .

# check if GITHUB_TOKEN is valid
check_github_token:
	@curl -s -H "Authorization: token $(GITHUB_TOKEN)" \
		https://api.github.com/user | jq '{login, name, type}'
	@curl -sI -H "Authorization: token $(GITHUB_TOKEN)" \
		https://api.github.com/user | grep -i x-oauth-scopes

doc: gen_files 
	@echo "*** Generating README.md with TOC ..."
	@touch $(README)
	@chmod 600 $(README)
	$(MARKDOWN_TOC_PROG) -i $(MAIN_MD) -o $(README) --glossary ${GLOSSARY_FILE} -f
	@chmod 444 $(README)
	$(MARKDOWN_TOC_PROG) -no-credit -i docs/ChangeLog.md -o ./ChangeLog.md --glossary ${GLOSSARY_FILE} -f

docs: doc

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
