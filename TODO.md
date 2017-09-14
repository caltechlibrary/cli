
# Action Items

## Bugs

## Next

+ [ ] add support for paged help (E.g. -help TERM would give you help page for TERM)
+ [ ] remove default prefix from cli.New(), i.e. prefix is always explicitly set

## Someday, Maybe

+ [ ] simplify pattern for rendering cfg.UsageText, cfg.DescriptionText, cfg.ExampleText
    + could change parameters for .New() to include explicit usage page sections
    + could provide a simple apply appName function  that would replace '%s' with appName
+ [ ] add multi-option definition support and remove dependency on native flag package

## Completed

+ [x] add method for common handling of input, output and Stderr
    + if filename specified the return file handle for it otherwise return handle for stdin/stdout/stderr as fallback
