
## EXAMPLE

```
    pkgassets MAP_VARAIBLE_NAME NAME_OF_DIRECTORY_HOLDING_ASSETS
```

This will result in a Go of type map[string][]byte holding the assets discovered by walking the directory
tree provided. The map's key will represent a path (beginning with "/") pointing at the asset ingested.

```shell
    pkgassets DefaultSite htdocs
```

Assuming that _htdocs_ held

+ index.html
+ css/site.css

In this example the htdocs directory will be crawled and all the files found harvested as a an asset. The
path in the map will not include htdocs and would result in a Go source file like

```golang
    package defaultsite

    var DefaultSite = map[string][]byte{
        "/index.html": []byte{}, // ... the contents of index.html would be here ...
        "/css/site.css": []byte{}, // ... the contents of css/site.css would be here ...
    }
```

If a package name is not provided then the package name will a lowercase name of the map variable name (e.g. 
"var DefaultSite" becomes "package defaultsite"). Likewise if a output name is not provided then the file
name will be the name of the package plus the ".go" extension.


