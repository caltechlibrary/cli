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
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//
// the following are pre v0.0.5 structs anf funcs, they are here for
// compatibility reasons until I convert the tools written for Caltech Library
// are converted to the new v0.0.5-dev approach.
//

// Config holds the merged environment and options selected when the program runs
// DEPRECIATED
type Config struct {
	appName         string            `json:"app_name"`
	version         string            `json:"version"`
	EnvPrefix       string            `json:"env_prefix"`
	LicenseText     string            `json:"license_text"`
	VersionText     string            `json:"version_text"`
	UsageText       string            `json:"usage_text"`
	DescriptionText string            `json:"description_text"`
	ExampleText     string            `json:"example_text"`
	OptionText      string            `json:"option_text"`
	topics          map[string]string `json:"help_pages"`
	examples        map[string]string `json:"examples"`

	// Options are the running options for an application, often this can be expressed as a cli
	// parameter or an environment variable.
	Options map[string]string `json:"options"`
}

// New returns an initialized Config structure
// DEPRECIATED
func New(appName, envPrefix, version string) *Config {
	prefix := strings.TrimSpace(envPrefix)
	return &Config{
		appName:         appName,
		version:         version,
		EnvPrefix:       prefix,
		LicenseText:     "",
		UsageText:       "",
		DescriptionText: "",
		OptionText:      "",
		ExampleText:     "",
		VersionText:     appName + " " + version,
		topics:          map[string]string{},
		examples:        map[string]string{},
		// Data used when processing options
		Options: make(map[string]string),
	}
}

// AddTopic takes a topic and text and address it to the help index
// DEPRECIATED
func (cfg *Config) AddHelp(topic, text string) {
	cfg.topics[topic] = text
}

// Help formats a help topics
// DEPRECIATED
func (cfg *Config) Help(topics ...string) string {
	text := []string{}
	notFound := false
	if len(topics) == 0 {
		return fmtTopics(fmt.Sprintf("Try '%s -help TOPIC' where TOPIC is one of ", cfg.appName), cfg.topics)
	}
	for _, topic := range topics {
		if pg, ok := cfg.topics[topic]; ok == true {
			text = append(text, pg)
		} else {
			text = append(text, fmt.Sprintf("%q not found", topic))
			notFound = true
		}
	}
	if notFound == true {
		text = append(text, fmtTopics(fmt.Sprintf("Try '%s -help TOPIC' where TOPIC is one of ", cfg.appName), cfg.topics))
	}
	return strings.Join(text, "") + "\n\n"
}

// AddExample takes a topic and example text and adds it to the examples index
// DEPRECIATED
func (cfg *Config) AddExample(topic, text string) {
	cfg.examples[topic] = text
}

// Example formats example topics
// DEPRECIATED
func (cfg *Config) Example(topics ...string) string {
	text := []string{}
	notFound := false
	if len(topics) == 0 {
		return fmtTopics(fmt.Sprintf("Try '%s -example TOPIC' where TOPIC is one of ", cfg.appName), cfg.examples)
	}
	for _, topic := range topics {
		if pg, ok := cfg.examples[topic]; ok == true {
			text = append(text, pg)
		} else {
			text = append(text, fmt.Sprintf("%q not found", topic))
			notFound = true
		}
	}
	if notFound == true {
		text = append(text, fmtTopics(fmt.Sprintf("Try '%s -help TOPIC' where TOPIC is one of ", cfg.appName), cfg.topics))
	}
	return strings.Join(text, "\n\n") + "\n\n"
}

// Usage returns a string describe how to use the cli, includes USAGE, SYNOPSIS, OPTIONS, etc.
// DEPRECIATED
func (cfg *Config) Usage() string {
	var text []string
	if len(cfg.UsageText) > 0 {
		text = append(text, cfg.UsageText)
	}
	if len(cfg.DescriptionText) > 0 {
		text = append(text, cfg.DescriptionText)
	}
	if len(cfg.OptionText) > 0 {
		text = append(text, cfg.OptionText)
	}

	// Loop through the flags describing cli options
	i := 0
	flag.VisitAll(func(f *flag.Flag) {
		text = append(text, "\t-"+f.Name+"\t"+f.Usage+"\n")
		i++
	})
	if i > 0 {
		text = append(text, "\n")
	}

	if len(cfg.ExampleText) > 0 {
		text = append(text, cfg.ExampleText)
	}

	// Display additional topics and examples if available
	if len(cfg.topics) > 0 {
		text = append(text, fmtTopics("Related topics: ", cfg.topics))
	}
	if len(cfg.examples) > 0 {
		text = append(text, fmtTopics("Related examples: ", cfg.examples))
	}
	if len(cfg.VersionText) > 0 {
		text = append(text, cfg.VersionText)
	}
	return strings.Join(text, "")
}

// License returns a string containing the application's license
// DEPRECIATED
func (cfg *Config) License() string {
	if len(cfg.LicenseText) > 0 {
		return cfg.LicenseText
	}
	return ""
}

// Version returns a string containing the application's version
// DEPRECIATED
func (cfg *Config) Version() string {
	if len(cfg.VersionText) > 0 {
		return cfg.VersionText
	}
	return ""
}

// Get returns the associate value of the key in the Config struct's Options
// DEPRECIATED
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
// DEPRECIATED
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
// DEPRECIATED
func (cfg *Config) MergeEnvBool(envVar string, flagValue bool) bool {
	s := os.Getenv(cfg.EnvPrefix + "_" + strings.ToUpper(envVar))
	envVal, _ := strconv.ParseBool(s)
	if envVal == true || flagValue == true {
		return true
	}
	return false
}

// CheckOption checks the trimmer string value, if len is 0, log an error message and if required is true exit(1)
// else return  the value passed in.
// DEPRECIATED
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
// DEPRECIATED
func (cfg *Config) StandardOptions(showHelp, showExamples, showLicense, showVersion bool, args []string) string {
	if showHelp == true {
		if len(args) > 0 {
			return cfg.Help(args...)
		} else {
			return cfg.Usage()
		}
	}
	if showExamples == true {
		if len(args) > 0 {
			return cfg.Example(args...)
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