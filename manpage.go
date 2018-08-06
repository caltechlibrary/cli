package cli

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

// md2man will try to do a crude conversion of Markdown
// to nroff man macros.
func md2man(src []byte) []byte {
	codeBlock := false
	lines := strings.Split(string(src), "\n")
	for i, line := range lines {
		if codeBlock == false {
			// Scan line for formatting conversions
		}
		// Scan line for code block handling
		if strings.HasPrefix(line, "```") {
			if codeBlock {
				lines[i] = `.EP`
				codeBlock = false
			} else {
				lines[i] = `.EX`
				codeBlock = true
			}
		}
	}
	return []byte(strings.Join(lines, "\n"))
}

// GenerateManPage writes a manual page suitable for running through
// `nroff --man`. May need some human clean up depending on content and
// internal formatting (e.g markdown style, spacing, etc.)
func (c *Cli) GenerateManPage(w io.Writer) {
	var parts []string

	// .TH {appName} {section_no} {version} {date}
	fmt.Fprintf(w, ".TH %s %d %q %q\n", c.appName, 1, time.Now().Format("2006 Jan 02"), strings.TrimSpace(strings.Replace(c.Version(), c.appName+" ", "", 1)))

	parts = append(parts, fmt.Sprintf(".TP\n\\fB%s\\fP", c.appName))
	if len(c.options) > 0 {
		parts = append(parts, "[OPTIONS]")
	}

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
	// .SH USAGE
	fmt.Fprintf(w, ".SH USAGE\n%s\n", strings.Join(parts, " "))

	// .SH SYNOPSIS
	// .SH DESCRIPTION
	if section, ok := c.Documentation["description"]; ok == true {
		fmt.Fprintf(w, ".SH SYNOPSIS\n%s\n", section)
	}

	if len(c.options) > 0 {
		fmt.Fprintf(w, ".SH OPTIONS\n")
		parts := []string{}
		if len(c.env) > 0 {
			parts = append(parts, ".TP\nOptions will override any corresponding environment settings.\n")
		}
		if len(c.actions) > 0 {
			parts = append(parts, ".TP\nOptions are shared between all actions and must precede the action on the command line.\n")
		}
		if len(parts) > 0 {
			fmt.Fprintf(w, "%s", strings.Join(parts, ""))
		}
		keys := []string{}
		for k, _ := range c.options {
			keys = append(keys, k)
		}
		// Sort the keys alphabetically and display output
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Fprintf(w, ".TP\n\\fB%s\\fP\n%s\n", k, c.options[k])
		}
	}

	if len(c.actions) > 0 {
		fmt.Fprintf(w, ".SS ACTIONS\n")
		keys := []string{}
		for k, _ := range c.actions {
			keys = append(keys, k)
		}
		// Sort the keys alphabetically and display output
		sort.Strings(keys)
		for _, k := range keys {
			usage := c.Action(k)
			fmt.Fprintf(w, ".TP\n\\fB%s\\fP\n%s\n", k, usage)
		}
	}

	if len(c.env) > 0 {
		fmt.Fprintf(w, ".SS ENVIRONMENT\n")
		if len(c.options) > 0 {
			fmt.Fprintf(w, "Environment variables can be overridden by corresponding options\n")
		}
		keys := []string{}
		for k, _ := range c.env {
			keys = append(keys, k)
		}
		// Sort the keys alphabetically and display output
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Fprintf(w, ".TP\n\\fB%s\\fP\n%s\n", k, c.env[k].Usage)
		}
	}

	// .SH EXAMPLES
	if section, ok := c.Documentation["examples"]; ok == true {
		//FIXME: Need to convert Markdown of examples into nroff with
		// with man macros.
		fmt.Fprintf(w, ".SH EXAMPLES\n")
		fmt.Fprintf(w, ".TP\n%s\n", md2man(section))
	}

	/*
		// .SH SEE ALSO
		fmt.Fprintf(w, ".SH SEE ALSO\n")
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
					links := []string{}
					for _, key := range keys {
						links = append(links, "%s", key)
					}
					fmt.Fprintf(w, ".SH SEE ALSO\n%s\n", strings.Join(links, ", "))
				}
			}
		// .BUGS
		fmt.Fprintf(w, ".SH BUGS\n")
		// AUTHORS
		fmt.Fprintf(w, ".SH AUTHORS\n")
		// COPYRIGHT
		fmt.Fprintf(w, ".SH COPYRIGHT\n")
	*/
}
