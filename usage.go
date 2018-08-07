//
// usage.go - handles rendering the "usage" message for cli package
//
package cli

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// Usage writes a help page to io.Writer provided. Documentation is based on
// the application's metadata like app name, version, options, actions, etc.
func (c *Cli) Usage(w io.Writer) {
	var parts []string
	parts = append(parts, c.appName)
	if len(c.options) > 0 {
		parts = append(parts, "[OPTIONS]")
	}
	// Add parts defined by AddParams()
	if len(c.actions) > 0 {
		if len(c.params) > 0 {
			parts = append(parts, c.params...)
		}
		if c.ActionsRequired {
			parts = append(parts, "ACTION [ACTION PARAMETERS...]")
		} else {
			parts = append(parts, "[ACTION] [ACTION PARAMETERS...]")
		}
	} else if len(c.params) > 0 {
		parts = append(parts, c.params...)
	}
	fmt.Fprintf(w, "\nUSAGE: %s\n\n", strings.Join(parts, " "))

	if section, ok := c.Documentation["description"]; ok == true {
		fmt.Fprintf(w, "SYNOPSIS\n\n%s\n\n", section)
	}

	if len(c.env) > 0 {
		fmt.Fprintf(w, "ENVIRONMENT\n\n")
		if len(c.options) > 0 {
			fmt.Fprintf(w, "Environment variables can be overridden by corresponding options\n\n")
		}
		keys := []string{}
		padding := 0
		for k, _ := range c.env {
			keys = append(keys, k)
			if len(k) > padding {
				padding = len(k) + 1
			}
		}
		// Sort the keys alphabetically and display output
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Fprintf(w, "    %s  %s\n", padRight(k, " ", padding), c.env[k].Usage)
		}
		fmt.Fprintf(w, "\n\n")
	}

	if len(c.options) > 0 {
		fmt.Fprintf(w, "OPTIONS\n\n")
		if len(c.env) > 0 {
			fmt.Fprintf(w, "Options will override any corresponding environment settings\n\n")
		}
		keys := []string{}
		padding := 0
		for k, _ := range c.options {
			keys = append(keys, k)
			if len(k) > padding {
				padding = len(k) + 1
			}
		}
		// Sort the keys alphabetically and display output
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Fprintf(w, "    %s  %s\n", padRight(k, " ", padding), c.options[k])
		}
		fmt.Fprintf(w, "\n\n")
	}

	if len(c.actions) > 0 {
		fmt.Fprintf(w, "ACTIONS\n\n")
		keys := []string{}
		padding := 0
		for k, _ := range c.actions {
			keys = append(keys, k)
			if len(k) > padding {
				padding = len(k) + 1
			}
		}
		// Sort the keys alphabetically and display output
		sort.Strings(keys)
		for _, k := range keys {
			usage := c.Action(k)
			fmt.Fprintf(w, "    %s  %s\n", padRight(k, " ", padding), usage)
		}
		fmt.Fprintf(w, "\n\n")
	}

	if section, ok := c.Documentation["examples"]; ok == true {
		fmt.Fprintf(w, "EXAMPLES\n\n%s\n\n", section)
	}

	if len(c.Documentation) > 0 {
		keys := []string{}
		for k, _ := range c.actions {
			if k != "description" && k != "examples" && k != "index" {
				keys = append(keys, k)
			}
		}
		if len(keys) > 0 {
			// Sort the keys alphabetically and display output
			sort.Strings(keys)
			fmt.Fprintf(w, "See %s -help TOPIC for topics - %s\n\n", c.appName, strings.Join(keys, ", "))
		}
	}

	fmt.Fprintf(w, "%s\n", c.version)
}
