

# cli: building better command line interfaces

_cli_ is a Golang package intended to encourage a more consistant 
command line user interface in programs written for Caltech Library
using Go. 

Features include:

+ Code generation for a command line tool
+ Integrated support for generating Markdown docs from a cli program
+ Integrated support for generating man pages from a cli program
+ Short/Long option (flag) support for both the command and verbs

## Command line tools

Two command line tools come with the _cli_ package.

+ [cligenerate](docs/cligenerate.html) - will generate a skelton command line program
+ [pkgassets](docs/pkgassets.html) - will take Markdown docs and create a go program file where each doc's name is the key and contents are the byte array value representing the docs


