//
// cligenerate is a cli application generator, e.g. a cmd/mything/main.go. It also demonstrates the use of the cli package.
//
// @author R. S. Doiel, <rsdoiel@library.caltech.edu>
//
// Copyright (c) 2021, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"fmt"
	"os"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
)

var (
	synopsis = `The cli package and cli application generator 
provides a standard way to construct a command line interfaces
encouraging uniformity across applications built at Caltech
Library. The _cligenerator_ program is intended primarily to be
an example of how to use the *cli* package.
`

	description = `This is a cli application generator. It also demonstrates
how to use the cli package.
`

	examples = `Example generating a new "helloworld"

` + "```" + `
    cligenerate -app=helloworld \
        -name="@author Jane Doe, <jane.doe@example.edu>" \
        -synopsis="This is a demo cli" \
		-use-description=README.md \
        -use-license=LICENSE
` + "```" + `

`

	bugs = `_cligenerator_ is only a proof of concept implementation`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	inputFName       string
	outputFName      string
	newLine          bool
	quiet            bool
	prettyPrint      bool
	generateMarkdown bool
	generateManPage  bool

	// Application Options
	appName             string
	appAuthor           string
	appSynopsis         string
	descriptionFilename string
	examplesFilename    string
	bugsFilename        string
	licenseFilename     string
)

func main() {
	app := cli.NewCli(cli.Version)

	// Add Help Docs
	app.SectionNo = 1 // Manual page section number to document
	app.AddHelp("synopsis", []byte(synopsis))
	app.AddHelp("description", []byte(description))
	app.AddHelp("examples", []byte(examples))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display examples")
	app.StringVar(&inputFName, "i,input", "", "input file name")
	app.StringVar(&outputFName, "o,output", "", "output file name")
	app.BoolVar(&newLine, "nl,newline", false, "if true add a trailing newline")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.BoolVar(&prettyPrint, "p,pretty", false, "pretty print output")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")

	// Application Options
	app.StringVar(&appName, "app", "[YOUR APP NAME GOES HERE]", "set the name of your generated app (e.g. helloworld)")
	app.StringVar(&appSynopsis, "synopsis", "[SHORT APP DESCRIPTION GOES HERE]", "set a short application synopsis (e.g. says 'Hello World!')")
	app.StringVar(&appAuthor, "name,author", "[YOUR AUTHOR STRING GOES HERE]", "set the author name (e.g. '@author Jane Doe, <jane.doe@example.edu>')")

	app.StringVar(&descriptionFilename, "use-description", "", "filename holding a detailed description of application.")
	app.StringVar(&examplesFilename, "use-examples", "", "filename holding examples")
	app.StringVar(&licenseFilename, "use-license", "LICENSE", "filename holding the license")
	app.StringVar(&bugsFilename, "use-bugs", "", "filename holding bugs")

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
	if generateMarkdown {
		app.GenerateMarkdown(app.Out)
		os.Exit(0)
	}
	if generateManPage {
		app.GenerateManPage(app.Out)
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

	// Run the app!
	srcCode := cli.Generate(appName, appSynopsis, appAuthor, descriptionFilename, examplesFilename, bugsFilename, licenseFilename)
	fmt.Fprintf(app.Out, "%s", srcCode)

	if newLine {
		fmt.Fprintln(app.Out, "")
	}
}
