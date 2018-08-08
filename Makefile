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
	go build -o bin/cligenerate$(EXT) cmd/cligenerate/cligenerate.go

install:
	env GOBIN=$(HOME)/bin go install cmd/cligenerate/cligenerate.go

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi

test:
	go test

man: build
	mkdir -p man/man1
	bin/cligenerate -generate-manpage | nroff -Tutf8 -man > man/man1/cligenerate.1

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

