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
	go install cmd/cligenerate/cligenerate.go

clean:
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d bin ]; then rm -fR bin; fi

test:
	go test

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

