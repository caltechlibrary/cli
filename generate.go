//
// generator.go - provides a quick and dirty skeleton for cli based apps.
//
package cli

import (
	"fmt"
	"strings"
)

var (
	commentBlock = `//
// %s %s
//
// %s
`
	importsBlock = `

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
)

`

	varBlock = `

var (
	description = %s

	examples = %s 

	// Standard Options
	showHelp bool
	showLicense bool
	showVersion bool
	showExamples bool
	inputFName string
	outputFName string
	newLine bool
	quiet bool
	prettyPrint bool
	generateMarkdownDocs bool

	// Application Options
)

`

	mainBlock = `

func main() {
	//FIXME: Replace with your base package .Version attribute
	app := cli.NewCli("v0.0.0")
	//FIXME: if you need the app name then...
	//appName := app.AppName()

	// Add Help Docs
	app.AddHelp("description", []byte(description))
	app.AddHelp("examples", []byte(examples))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "e,examples", false, "display examples")
	app.StringVar(&inputFName, "i,input", "", "input file name")
	app.StringVar(&outputFName, "o,output", "", "output file name")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&prettyPrint, "p,pretty", false, "pretty print output")
	app.BoolVar(&generateMarkdownDocs, "generate-markdown-docs", false, "output documentation in Markdown")

	// Application Options
	//FIXME: Add any application specific options

	// Action verbs (e.g. app.AddAction(STRING_VERB, FUNC_POINTER, STRING_DESCRIPTION)
	//FIXME: If the application is verb based add your verbs here

	// We're ready to process args
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error

	app.Eout = os.Stderr
	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Handle options
	if generateMarkdownDocs {
		app.GenerateMarkdownDocs(app.Out)
		os.Exit(0)
	}
	if showHelp || showExamples {
		if len(args) > 0 {
			fmt.Fprintf(app.Out, app.Help(args...))
		} else {
		 	app.Usage(app.Out)
		}
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintln(app.Out, app.License())
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintln(app.Out, app.Version())
		os.Exit(0)
	}

	// FIXME: Application Option Parsing

	// FIXME: Application running code, e.g.
	// Run the app!
	//os.Exit(app.Run(args))

	if newLine {
		fmt.Fprintln(app.Out, "")
	}
}

`
)

func backQuote(s string) string {
	return "`" + s + "\n`"
}

// Generate creates main.go source code
func Generate(appName, description, author, license string) []byte {
	blocks := []string{}
	// Setup the initial comment block
	blocks = append(blocks, fmt.Sprintf(commentBlock, appName, description, author))
	// Add the license
	blocks = append(blocks, fmt.Sprintf("//%s\n", strings.Replace(license, "\n", "\n// ", -1)))
	// Add Package name
	blocks = append(blocks, "package main")
	// Add Standard Imports
	blocks = append(blocks, importsBlock)
	// Add Global vars
	blocks = append(blocks, fmt.Sprintf(varBlock, backQuote(description), backQuote("[FIXME: examples go here]")))
	// Add Main
	blocks = append(blocks, mainBlock)

	// Convert to byte array and return
	return []byte(strings.Join(blocks, ""))
}