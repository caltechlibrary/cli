#
# Simple Makefile
#
PROJECT = cli

VERSION = $(shell grep -m1 "Version = " $(PROJECT).go | cut -d\` -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\   -f 2)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
    EXT = .exe
endif

build: 
	go build -o bin/pkgassets$(EXT) cmd/pkgassets/pkgassets.go 
	go build -o bin/cligenerate$(EXT) cmd/cligenerate/cligenerate.go


man: build
	mkdir -p man/man1
	bin/cligenerate -generate-manpage | nroff -Tutf8 -man > man/man1/cligenerate.1
	bin/pkgassets -generate-manpage | nroff -Tutf8 -man > man/man1/pkgassets.1


test:
	go test
	cd pkgassets && go test

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi


install: build
	env GOBIN=$(GOPATH)/bin go install cmd/cligenerate/cligenerate.go
	env GOBIN=$(GOPATH)/bin go install cmd/pkgassets/pkgassets.go


dist/linux-amd64:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/cligenerate cmd/cligenerate/cligenerate.go
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/pkgassets cmd/pkgassets/pkgassets.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/cligenerate.exe cmd/cligenerate/cligenerate.go
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/pkgassets.exe cmd/pkgassets/pkgassets.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macos-amd64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/cligenerate cmd/cligenerate/cligenerate.go
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/pkgassets cmd/pkgassets/pkgassets.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macos-arm64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/cligenerate cmd/cligenerate/cligenerate.go
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/pkgassets cmd/pkgassets/pkgassets.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-arm64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/cligenerate cmd/cligenerate/cligenerate.go
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/pkgassets cmd/pkgassets/pkgassets.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-amd7.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	./package-versions.bash > dist/package-versions.txt

release: bootstrap distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macos-amd64 dist/macos-arm64 dist/raspbian-arm7

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash

