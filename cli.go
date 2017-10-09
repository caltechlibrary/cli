/**
 * cli is a package intended to encourage some standardization in the command line user interface for programs
 * developed for Caltech Library.
 *
 * @author R. S. Doiel, <rsdoiel@caltech.edu>
 *
 * Copyright (c) 2016, Caltech
 * All rights not granted herein are expressly reserved by Caltech.
 *
 * Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 *
 * 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */
package cli

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const Version = `v0.0.2`

// Config holds the merged environment and options selected when the program runs
type Config struct {
	appName         string            `json:"app_name"`
	version         string            `json:"version"`
	EnvPrefix       string            `json:"env_prefix"`
	LicenseText     string            `json:"license_text"`
	VersionText     string            `json:"version_text"`
	UsageText       string            `json:"usage_text"`
	DescriptionText string            `json:"description_text"`
	OptionsLabel    string            `json:"option_label"`
	topics          map[string]string `json:"help_pages"`
	examples        map[string]string `json:"examples"`
	// Options are the running options for an application, often this can be expressed as a cli
	// parameter or an environment variable.
	Options map[string]string `json:"options"`
}

// New returns an initialized Config structure
func New(appName, envPrefix, license, version string) *Config {
	prefix := strings.TrimSpace(envPrefix)
	return &Config{
		appName:         appName,
		version:         version,
		EnvPrefix:       prefix,
		LicenseText:     license,
		UsageText:       "",
		DescriptionText: "",
		OptionsLabel:    "OPTIONS",
		VersionText:     appName + " " + version,
		topics:          make(map[string]string),
		examples:        make(map[string]string),
		Options:         make(map[string]string),
	}
}

// FmtAppName takes an appName and text replacing each occurrence of %s
// with the value of AppName.
func (cfg *Config) FmtAppName(appName, text string) string {
	rtext := []interface{}{}
	c := strings.Count(text, "%s")
	for i := 0; i < c; i++ {
		rtext = append(rtext, appName)
	}
	return fmt.Sprintf(text, rtext...)
}

// AddTopic takes a topic and text and address it to the help index
func (cfg *Config) AddHelp(topic, text string) {
	cfg.topics[topic] = text
}

// AddExample takes a topic and example text and adds it to the examples index
func (cfg *Config) AddExample(topic, text string) {
	cfg.examples[topic] = text
}

func (cfg *Config) Usage() string {
	var text []string
	if len(cfg.UsageText) > 0 {
		text = append(text, cfg.UsageText+"\n\n")
	}
	if len(cfg.DescriptionText) > 0 {
		text = append(text, cfg.DescriptionText+"\n\n")
	}
	if len(cfg.OptionsLabel) > 0 {
		text = append(text, cfg.OptionsLabel+"\n\n")
	}
	// Loop through the flags
	i := 0
	flag.VisitAll(func(f *flag.Flag) {
		text = append(text, "\t-"+f.Name+"\t"+f.Usage+"\n")
		i++
	})
	if i > 0 {
		text = append(text, "\n")
	}
	if len(cfg.topics) > 0 {
		// Sort the topics keys
		keys := []string{}
		for k, _ := range cfg.topics {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		text = append(text, "Topics (e.g. -help KEYWORD)\n\n")
		for _, k := range keys {
			text = append(text, "+ "+k+"\n")
		}
		text = append(text, "\n")
	}
	if len(cfg.examples) > 0 {
		// Sort the example keys
		keys := []string{}
		for k, _ := range cfg.examples {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		// For each sorted key add it to page.
		text = append(text, "Examples (e.g. -example KEYWORD)\n\n")
		for _, k := range keys {
			text = append(text, "+ "+k+"\n")
		}
		text = append(text, "\n")
	}
	if len(cfg.VersionText) > 0 {
		text = append(text, cfg.VersionText)
	}
	return strings.Join(text, "")
}

func (cfg *Config) License() string {
	if len(cfg.LicenseText) > 0 {
		return cfg.LicenseText
	}
	return ""
}

func (cfg *Config) Version() string {
	if len(cfg.VersionText) > 0 {
		return cfg.VersionText
	}
	return ""
}

func (cfg *Config) Get(key string) string {
	if val, ok := cfg.Options[key]; ok == true {
		return val
	}
	return ""
}

// MergeEnv merge environment variables into the configuration structure.
// options are
// + prefix - e.g. EPGO, name space before the first underscore in the envinronment
//   + prefix plus uppercase key forms the complete environment variable name
// + key - the field map (e.g ApiURL maps to API_URL in EPGO_API_URL for prefix EPGO)
// + proposedValue - the proposed value, usually the value from the flags passed in (an empty string means no value provided)
//
// returns the new value of the environment string merged
func (cfg *Config) MergeEnv(envVar, flagValue string) string {
	val := strings.TrimSpace(flagValue)
	if len(val) > 0 {
		cfg.Options[envVar] = val
	} else {
		cfg.Options[envVar] = os.Getenv(cfg.EnvPrefix + "_" + strings.ToUpper(envVar))
	}
	return cfg.Options[envVar]
}

// MergeEnvBool till pick use flagValue if present, otherwise is the environment value.
// It returns the value selected (e.g. useful when combined with CheckOption()
func (cfg *Config) MergeEnvBool(envVar string, flagValue bool) bool {
	envVal := (func(envVar string) bool {
		s := os.Getenv(cfg.EnvPrefix + "_" + strings.ToUpper(envVar))
		switch strings.TrimSpace(strings.ToLower(s)) {
		case "true":
			return true
		case "t":
			return true
		case "1":
			return true
		default:
			return false
		}
	}(envVar))

	if envVal == true || flagValue == true {
		return true
	}
	return false
}

// CheckOption checks the trimmer string value, if len is 0, log an error message and if required is true exit(1)
// else return  the value passed in.
func (cfg *Config) CheckOption(envVar, value string, required bool) string {
	value = strings.TrimSpace(value)
	if len(value) == 0 && required == true {
		log.Printf("Missing %s_%s", strings.ToUpper(cfg.EnvPrefix), strings.ToUpper(envVar))
		os.Exit(1)
	}
	return value
}

// StandardOptions() processing the booleans associated with standard options and
// any additional cli parameters in args return text as a string.
func (cfg *Config) StandardOptions(showHelp, showExamples, showLicense, showVersion bool, args []string) string {
	if showHelp == true {
		if len(args) > 0 {
			text := []string{}
			for _, arg := range args {
				if pg, ok := cfg.topics[arg]; ok == true {
					text = append(text, pg+"\n\n")
				} else {
					text = append(text, "No information for "+arg+"\n\n")
				}
			}
			return strings.Join(text, "")
		} else {
			return cfg.Usage()
		}
	}
	if showExamples == true {
		if len(args) > 0 {
			text := []string{}
			for _, arg := range args {
				if pg, ok := cfg.examples[arg]; ok == true {
					text = append(text, pg+"\n\n")
				} else {
					text = append(text, "No example for "+arg+"\n\n")
				}
			}
			return strings.Join(text, "")
		} else {
			return cfg.Usage()
		}
	}
	if showLicense == true {
		return cfg.License()
	}
	if showVersion == true {
		return cfg.Version()
	}
	return ""
}

// Open accepts a filename, fallbackFile (usually os.Stdout, os.Stdin, os.Stderr) and returns
// a file pointer and error.  It is a conviences function for wrapping stdin, stdout, stderr
// If filename is "-" or filename is "" then fallbackFile is used.
func Open(filename string, fallbackFile *os.File) (*os.File, error) {
	if len(filename) == 0 || filename == "-" {
		return fallbackFile, nil
	}
	return os.Open(filename)
}

// Create accepts a filename, fallbackFile (usually os.Stdout, os.Stdin, os.Stderr) and returns
// a file pointer and error.  It is a conviences function for wrapping stdin, stdout, stderr
// If filename is "-" or filename is "" then fallbackFile is used.
func Create(filename string, fallbackFile *os.File) (*os.File, error) {
	if len(filename) == 0 || filename == "-" {
		return fallbackFile, nil
	}
	return os.Create(filename)
}

// CloseFile accepts a filename and os.File pointer, if filename is "" or "-" it skips the close
// otherwise is does a fp.Close() on the file.
func CloseFile(filename string, fp *os.File) error {
	if len(filename) == 0 || filename == "-" {
		return nil
	}
	return fp.Close()
}

// ReadLines accepts a file pointer to read and returns an array of lines.
func ReadLines(in *os.File) ([]string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err := scanner.Err()
	return lines, err
}

// IsPipe accepts a file pointer and returns true if data on file pointer is
// from a pipe or false if not.
func IsPipe(in *os.File) bool {
	finfo, err := in.Stat()
	if err == nil && (finfo.Mode()&os.ModeCharDevice) == 0 {
		return true
	}
	return false
}
