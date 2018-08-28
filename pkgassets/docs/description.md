
## Description

_pkgassets_ generates a Go source directory whos file assets are embedded in a `map[string][]byte` variable. 
This is useful where you want to embed web content, template source code, help docs and other assets that 
can be used for default behavior in a Go command line program or service. 

The map content is harvested from directory holding the assets to be embedded. By default the
path key starts with a slash and does not include the hosting directory (e.g. htdocs/index.html 
would become /index.html if htdocs was used to harvest assets). The prefix and suffix on the
key can be modified based on _pkgassets_'s command line options.

