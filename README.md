

# package cli

_cli_ is a Golang package to encourage a more consistant command line user interface
in programs written for Caltech Library.

Features include:

+ A common `Config` object generated with `cli.New()`
+ A common `Create`, `Open` and `Close` file wrapper for integrating standard in/out/err as fallback
+ A `Readlines()` and `IsPipe()` for handling content read when working with text files from disc or standard in
+ The `Config` type supports the following functions
    + `AddHelp(topic, text string)` for adding help topics by page name
    + `Help(topics ...string)` string for searching your help topics by page name
    + `AddExample(topic, text string)` for adding example topics by page name
    + `Example(topics ...string) string` for searching your examples by page name
    + `Usage() string` builds a page for show general help - usage, description, options and examples
    + `License() string` builds a license page as a string
    + `Version() string` builds a version string
    + `Get(key string) string` gets an option discovered in the environment
    + `MergeEnv(envVar, flagValue string) string` merges the environment with default a string value
    + `MergeEnvBool(envVar string, flagValue bool) bool`  merges the environment with a default bool value
    + `CheckOption(envVar, value string, required bool) string` Checks the environment, applies a default and halts if error if needed
    + `StandardOptions(showHelp, showExamples, showLicense, showVersion bool, args []string) string` built-in process for help, version, etc.

