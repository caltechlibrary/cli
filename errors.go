package cli

import (
	"fmt"
	"os"
)

// ExitOnError is used by the cli programs to
// handle exit cuasing errors constitantly.
// E.g. it respects the -quiet flag past to it.
func ExitError(out *os.File, err error, quiet bool) {
	if err != nil {
		if quiet == false {
			fmt.Fprint(out, "%s\n", err)
		}
		os.Exit(1)
	}
}
