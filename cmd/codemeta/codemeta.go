//
// codemeta This program generates a codemeta JSON file and provides a means of updating an existing one
//
// @author R. S. Doiel, <rsdoiel@library.caltech.edu>
//
// Copyright (c) 2018, Caltech
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
	"io"
	"os"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/cli/codemeta"
)

var (
	synopsis = `This program generates/updates 'codemeta.json' files.`

	description = `
_codemeta_ generates or updates a [codemeta.json]() file documenting
a software or data project.
`

	examples = `
Generate a new empty codemeta.json file
(by default is generates it in the current
directory).

` + "```" + `
	codemeta create \
	  who "J. Doe" \
	  what "J's little helper"
	  version "0.0.0-alpha"
` + "```" + `

Update "janesapp"'s codemeta.json with specific
values.

` + "```" + `
    codemeta update who "Doe, Jane"
	codemeta update email "<jane.doe@bigco.example.com>"
	codemeta update what "Jane's program. It will be self-aware soon"
	codemeta update when "2018-08-28"
	codemeta update where "Freedonia"
	codemeta update use-license LICENSE
	codemeta update version "10.10.10"
` + "```" + `

Get the value of "version" from a codemeta.json file.

` + "```" + `
    codemeta get version
` + "```" + `

`

	bugs = `
codemeta will read an existing [codemeta.json]()
file updates using common terminalogy crossing walking
when known (e.g. namaste who to author[0].name).
If you need more nuanced control you're better 
off with a JSON editor.
`

	license = `
Copyright (c) 2018, Caltech
All rights not granted herein are expressly reserved by Caltech.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

`

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
)

// create - creates a 'codemeta.json' file
func create(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	fmt.Fprintf(eout, "create not implemented\n")
	return 1
}

// update - updates a 'codemeta.json' file
func update(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	fmt.Fprintf(eout, "update not implemented\n")
	return 1
}

// get - retrieves a value from a 'codemeta.json' file
func get(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	fmt.Fprintf(eout, "get not implemented\n")
	return 1
}

func main() {
	app := cli.NewCli(codemeta.Version)

	// Add Help Docs
	app.SectionNo = 1 // The manual page section number
	app.AddHelp("synopsis", []byte(synopsis))
	app.AddHelp("description", []byte(description))
	app.AddHelp("examples", []byte(examples))
	app.AddHelp("bugs", []byte(bugs))

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
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate Markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "output manpage markup")

	// Application Options
	//FIXME: Add any application specific options

	// Application Verbs
	vCreate := app.NewVerb("create", "creates a 'codemeta.json' file", create)
	vCreate.AddParams("TAG", "VALUE", "[TAG VALUE]", "...")

	vUpdate := app.NewVerb("update", "updates a 'codemeta.json' file", update)
	vUpdate.AddParams("TAG", "VALUE", "[TAG VALUE]", "...")

	vGet := app.NewVerb("get", "retrieves a value from the codemeta.json file", get)
	vGet.AddParams("TAG")

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

	// Application Logic
	switch {
	case len(args) == 0:
		app.Usage(app.Eout)
		os.Exit(1)
	case args[0] == "create":
	case args[0] == "update":
	case args[0] == "get":
	default:
		fmt.Fprintf(app.Eout, "Unknown action %q\n", args[0])
		os.Exit(1)
	}

	if newLine {
		fmt.Fprintln(app.Out, "")
	}
}
