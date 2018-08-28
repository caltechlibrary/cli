
# USAGE

	codemeta [OPTIONS]

## SYNOPSIS

This program generates/updates 'codemeta.json' files.

## DESCRIPTION


_codemeta_ generates or updates a [codemeta.json]() file documenting
a software or data project.


## OPTIONS

Below are a set of options available.

```
    -examples            display examples
    -generate-manpage    output manpage markup
    -generate-markdown   generate Markdown documentation
    -h, -help            display help
    -i, -input           input file name
    -l, -license         display license
    -nl, -newline        if true add a trailing newline
    -o, -output          output file name
    -p, -pretty          pretty print output
    -quiet               suppress error messages
    -v, -version         display version
```


## EXAMPLES


Generate a new empty codemeta.json file
(by default is generates it in the current
directory).

```
	codemeta create \
	  who "J. Doe" \
	  what "J's little helper"
	  version "0.0.0-alpha"
```

Update "janesapp"'s codemeta.json with specific
values.

```
    codemeta update who "Doe, Jane"
	codemeta update email "<jane.doe@bigco.example.com>"
	codemeta update what "Jane's program. It will be self-aware soon"
	codemeta update when "2018-08-28"
	codemeta update where "Freedonia"
	codemeta update use-license LICENSE
	codemeta update version "10.10.10"
```

Get the value of "version" from a codemeta.json file.

```
    codemeta get version
```



codemeta v0.0.15
