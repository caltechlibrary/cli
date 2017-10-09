

# package cli

_cli_ is a Golang package to encourage a more consistant command line user interface
in programs written for Caltech Library.

Features include:

+ Supports a common configuration for handling base options (e.g. -h, -l, -v for help, license and version)
+ FmtAppName() - formats a text with '%s' replacing it with the supplied app name
+ AddHelp() - adds a topic and related help page
+ AddExample() - adds a topic and related example
+ StandardOptions() - handle the typical -help, -example, -license, -version option booleans
+ MergeEnv() - a function to merge environment variables and command line options
+ Usage() - display the usual help page expected with the "-h, --help" option



