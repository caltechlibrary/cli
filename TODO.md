
# Action Items

## Bugs

## Next

+ [ ] See ...


## Someday, Maybe

+ [ ] a cli.GenerateCmd would create a skeleton app based on options and exported items in package
+ [ ] Need to word wrap output of cli help, usage, etc. based on console width and height

## Completed

+ [x] Environment processing stil doesn't feel right
    + Maybe have an EnvTYPE functions that adds the documentation and returns the value found or default, this would work like flag.StringVar() 
+ [x] add a way to format AppName into a text (e.g. DescriptionText, UsageText)
+ [x] add support for topic driven help pages (E.g. -help KEYWORD would give you help page for KEYWORD)
+ [x] add support for topic driven examples pages (E.g. -example KEYWORD would give you example page for KEYWORD)
+ [x] add StandardOptions() for handling showHelp, showExamples, showLicense, and showVersion
+ [x] add a way of adding topic drive help and example pages 
+ [x] remove default prefix from cli.New(), i.e. prefix should always explicitly set
+ [x] add method for common handling of input, output and Stderr
    + if filename specified the return file handle for it otherwise return handle for stdin/stdout/stderr as fallback
+ [x] Develop cli.App model that works for various tools
    + In, Out, Err for stdin, stdout, stderr as defined by app
         + In, Out, Err could be io.Reader, io.Writer, io.Writer interfaces
    + Documentation map for docs
    + Vocabulary that maps functions to parameters like git, dataset commands
    + Auto generate a usage statement based on Vocabulary and flags
    + Set License, Description texts
    + Links topic pages (e.g. examples and detail explanations)
    + Should beable to generate a cli skeleton easily for building cli
+ [x] simplify pattern for rendering cfg.UsageText, cfg.DescriptionText, cfg.ExampleText
    + could change parameters for .New() to include explicit usage page sections
    + could provide a simple apply appName function  that would replace '%s' with appName
+ [x] add multi-option definition support and remove dependency on native flag package
