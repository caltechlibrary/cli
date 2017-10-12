
# Action Items

## Bugs

## Next


## Someday, Maybe

+ [ ] add a way to format AppName into a text (e.g. DescriptionText, UsageText)
+ [ ] simplify pattern for rendering cfg.UsageText, cfg.DescriptionText, cfg.ExampleText
    + could change parameters for .New() to include explicit usage page sections
    + could provide a simple apply appName function  that would replace '%s' with appName
+ [ ] add multi-option definition support and remove dependency on native flag package

## Completed

+ [x] add support for topic driven help pages (E.g. -help KEYWORD would give you help page for KEYWORD)
+ [x] add support for topic driven examples pages (E.g. -example KEYWORD would give you example page for KEYWORD)
+ [x] add StandardOptions() for handling showHelp, showExamples, showLicense, and showVersion
+ [x] add a way of adding topic drive help and example pages 
+ [x] remove default prefix from cli.New(), i.e. prefix should always explicitly set
+ [x] add method for common handling of input, output and Stderr
    + if filename specified the return file handle for it otherwise return handle for stdin/stdout/stderr as fallback
