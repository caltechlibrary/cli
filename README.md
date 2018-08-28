

# cli: building better command line interfaces

_cli_ is a Golang package intended to encourage a more consistant 
command line user interface in programs written for Caltech Library
using Go. 

Features include:

+ A common `Cli` object generated with `cli.NewCli()` representing configuration, environment and command line
    + environment and options share the same approach as Go's standard flag package
    + options support multiple options (long/short version) in option's flag string separating options by comma
    + generate help and usage, sysnopsis, options, examples, license based on programaticly
+ A common `Create`, `Open` and `Close` file wrapper for integrating standard in/out/err as fallback
+ The `Cli` object supports the following functions
    + `AppName()` returns the compiled application name
    + `AddParams()` for documenting required and optional non-option parameters
    + `AddAction()` for verb style command lines, associates verb, func and doc string
    + `AddVerb()` for verb style command lines, associates verb and doc string
    + `AddFlagSet()` for associating flags with the command or verb
    + `AddHelp()` for adding help topics by page name
    + `Help()` string for searching your help topics by page name
    + `Usage() string` builds a page for show general help - synopsis, description, options and examples
    + `License() string` builds a license page as a string
    + `Version() string` builds a version string
    + `Run()` for running defined actions and returning an exit code suitable for passing to `os.Exit()`
    + `GenerateMarkdownDocs()` for generating documentation based on how the program in implemented
    + `GenerateManPage()` for generating nroff man page documentation

## Command line tools

Two command line tools come with the [cli](./) package.

+ [cligenerate](docs/cligenerate.html) - will generate a skelton command line program
+ [pkgassets](docs/pkgassets.html) - will take Markdown docs and create a go program file where each doc's name is the key and contents are the byte array value representing the docs


