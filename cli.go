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
	"log"
	"os"
	"strings"
)

// Config holds the merged environment and options selected when the program runs
type Config struct {
	appName         string            `json:"app_name"`
	version         string            `json:"version"`
	EnvPrefix       string            `json:"env_prefix"`
	LicenseText     string            `json:"license_text"`
	VersionText     string            `json:"version_text"`
	UsageText       string            `json:"usage_text"`
	DescriptionText string            `json:"description_text"`
	OptionsText     string            `json:"option_text"`
	ExampleText     string            `json:"example_text"`
	Options         map[string]string `json:"options"`
}

// New returns an initialized Config structure
func New(appName, envPrefix, license, version string) *Config {
	if envPrefix == "" {
		envPrefix = strings.ToUpper(appName)
	}
	return &Config{
		appName:         appName,
		version:         version,
		EnvPrefix:       strings.ToUpper(envPrefix),
		LicenseText:     license,
		UsageText:       "",
		DescriptionText: "",
		OptionsText:     "",
		ExampleText:     "",
		VersionText:     appName + " " + version,
		Options:         make(map[string]string),
	}
}

func (cfg *Config) Usage() string {
	var text []string
	if len(cfg.UsageText) > 0 {
		text = append(text, cfg.UsageText)
	}
	if len(cfg.DescriptionText) > 0 {
		text = append(text, cfg.DescriptionText)
	}
	if len(cfg.OptionsText) > 0 {
		text = append(text, cfg.OptionsText)
	}
	// Loop through the flags
	flag.VisitAll(func(f *flag.Flag) {
		text = append(text, "\t-"+f.Name+"\t"+f.Usage)
	})
	if len(cfg.ExampleText) > 0 {
		text = append(text, cfg.ExampleText)
	}
	text = append(text, "\n")
	if len(cfg.VersionText) > 0 {
		text = append(text, cfg.VersionText)
	}
	return strings.Join(text, "\n")
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
	if len(value) == 0 {
		log.Printf("Missing %s_%s", strings.ToUpper(cfg.appName), strings.ToUpper(envVar))
		if required == true {
			os.Exit(1)
		}
	}
	return value
}
