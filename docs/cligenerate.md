
# USAGE

	cligenerate [OPTIONS]

## SYNOPSIS

The cli package and cli application generator 
provides a standard way to construct a command line interfaces
encouraging uniformity across applications built at Caltech
Library. The _cligenerator_ program is intended primarily to be
an example of how to use the *cli* package.


## DESCRIPTION

This is a cli application generator. It also demonstrates
how to use the cli package.


## OPTIONS

Below are a set of options available.

```
    -app                set the name of your generated app (e.g. helloworld)
    -e, -examples       display examples
    -generate-manpage   generate man page
    -generate-markdown  generate markdown documentation
    -h, -help           display help
    -i, -input          input file name
    -l, -license        display license
    -name, -author      set the author name (e.g. '@author Jane Doe, <jane.doe@example.edu>')
    -nl, -newline       if true add a trailing newline
    -o, -output         output file name
    -p, -pretty         pretty print output
    -quiet              suppress error messages
    -synopsis           set a short application description (e.g. says 'Hello World!')
    -use-bugs           filename holding bugs
    -use-description    filename holding a detailed description of application.
    -use-examples       filename holding examples
    -use-license        filename holding the license
    -v, -version        display version
```


## EXAMPLES

Example generating a new "helloworld"

```
    cligenerate -app=helloworld \
        -name="@author Jane Doe, <jane.doe@example.edu>" \
        -decription="This is a demo cli" \
        -use-license=LICENSE
```


cligenerate v0.0.14
