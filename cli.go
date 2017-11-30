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
	"io"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

const Version = `v0.0.5`

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

// PopArg takes an array of strings and if array is not empty returns a string and the rest of the args.
// If array is empty it returns an empty string, when there are no more args it returns nil for the
// arg parameter
func PopArg(args []string) (string, []string) {
	var s string

	if args != nil && len(args) >= 1 {
		s = args[0]
		if len(args) > 1 {
			args = args[1:]
		} else {
			args = nil
		}
	}
	return s, args
}

//
// v0.0.5 brings a more wholistic approach to building a cli, not just configuring one.
// Below are the new structs and associated methods that will replace the original Config approach
// at some point.
//

// Action describes an "action" that a cli might take. Actions aren't prefixed with a "-".
type Action struct {
	// Name is usually a verb like list, test, build as needed by the cli
	Name string
	// Fn is action that will be run by Cli.Run() if Name is the first non-option arg on the command line
	Fn func(io.Reader, io.Writer, io.Writer, []string) int
	// Usage is a short description of what the action does and description of any expected additoinal parameters
	Usage string
}

// EnvAttribute describes expected environmental attributes associated with the cli app.
// It also provides the default value of the environmental attribute if missing from the environment.
type EnvAttribute struct {
	// Name is the environment variable (e.g. DATASET, USERNAME)
	Name string
	// Type holds the type name of the attribute, e.g. int, int64, float64, string, bool, uint, uint64, time.Duration
	Type string
	// BoolValue holds the default boolean
	BoolValue bool
	// IntValue holds the default int
	IntValue int
	// Int64Value holds the default int64
	Int64Value int64
	// UintValue holds the default uint
	UintValue uint
	// Uint64Value holds the default uint64
	Uint64Value uint64
	// Float64Value holds the default float64
	Float64Value float64
	// Dura1tionValue holds the default time.Duration
	DurationValue time.Duration
	// StringValue holds the default string
	StringValue string
	// Usage describes the environment variable role and expected setting
	Usage string
}

// Cli models the metadata for running a common cli program
type Cli struct {
	// In is usually set to os.Stdin
	In *os.File
	// Out is usually set to os.Stdout
	Out *os.File
	// Eout is usually set to os.Stderr
	Eout *os.File
	// Documentation specific help pages, e.g. -help example1
	Documentation map[string][]byte

	// application name based on os.Args[0]
	appName string
	// application version based on string passed in New
	version string
	// expected environmental variables used by app
	env map[string]*EnvAttribute
	// description of additoinal command line parameters
	params []string
	// description of short/long options and their doc strings
	options map[string]string
	// non-flag options, e.g. in the command line "go test", "test" would be the action string. Any
	// additional parameters would be handed of the associated Action.
	actions map[string]*Action
}

// NewCli creates an Cli instance, an Cli describes the running of the command line interface
// making it easy to expose the functionality in packages as command line tools.
func NewCli(version string) *Cli {
	appName := path.Base(os.Args[0])
	env := make(map[string]*EnvAttribute)
	options := make(map[string]string)
	actions := make(map[string]*Action)
	documentation := make(map[string][]byte)
	return &Cli{
		In:            os.Stdin,
		Out:           os.Stdout,
		Eout:          os.Stderr,
		Documentation: documentation,
		appName:       appName,
		version:       fmt.Sprintf("%s %s", appName, version),
		env:           env,
		params:        []string{},
		options:       options,
		actions:       actions,
	}
}

// AppName returns the application name as a string
func (c *Cli) AppName() string {
	return c.appName
}

// AddHelp takes a string keyword and byte slice of content and
// updates the Documentation attribute
func (c *Cli) AddHelp(keyword string, usage []byte) error {
	c.Documentation[keyword] = usage
	_, ok := c.Documentation[keyword]
	if ok == false {
		return fmt.Errorf("could not add help for %q", keyword)
	}
	return nil
}

// Help returns documentation on a topic
func (c *Cli) Help(keywords ...string) string {
	var sections []string

	for _, keyword := range keywords {
		usage, ok := c.Documentation[keyword]
		if ok == false {
			sections = append(sections, fmt.Sprintf("%q not documented", keyword))
			continue
		}
		sections = append(sections, fmt.Sprintf("%s\n\n%s", strings.ToUpper(keyword), usage))
	}
	return strings.Join(sections, "\n\n")
}

func splitOps(names string) []string {
	var ops []string

	switch {
	case strings.Contains(names, ","):
		// If we have a comma separated list of options, e.g. "h,help" for -h, -help
		ops = strings.Split(names, ",")
	case strings.Contains(names, " "):
		// If we have a space separated list of options
		ops = strings.Split(names, ",")
	default:
		ops = []string{names}
	}
	// clean up spaces and dash prefixes
	for i := 0; i < len(ops); i++ {
		op := strings.Trim(ops[i], "- ")
		ops[i] = op
	}
	return ops
}

func opsLabel(ops []string) string {
	parts := []string{}
	for _, op := range ops {
		parts = append(parts, "-"+op)
	}
	return strings.Join(parts, ", ")
}

// BoolVar updates c.options doc strings, then splits options and calls flag.BoolVar()
func (c *Cli) BoolVar(p *bool, names string, value bool, usage string) {
	// Prep to hand off to the flag package
	ops := splitOps(names)
	// Save for our internal option documentation
	label := opsLabel(ops)
	c.options[label] = usage
	// process with flag package
	for _, op := range ops {
		flag.BoolVar(p, op, value, usage)
	}
}

// IntVar updates c.options doc strings, then splits options and calls flag.IntVar()
func (c *Cli) IntVar(p *int, names string, value int, usage string) {
	// Prep to hand off to the flag package
	ops := splitOps(names)
	// Save for our internal option documentation
	label := opsLabel(ops)
	c.options[label] = usage
	// process with flag package
	for _, op := range ops {
		op = strings.TrimSpace(op)
		flag.IntVar(p, op, value, usage)
	}
}

// Int64Var updates c.options doc strings, then splits options and calls flag.Int64Var()
func (c *Cli) Int64Var(p *int64, names string, value int64, usage string) {
	// Prep to hand off to the flag package
	ops := splitOps(names)
	// Save for our internal option documentation
	label := opsLabel(ops)
	c.options[label] = usage
	// process with flag package
	for _, op := range ops {
		op = strings.TrimSpace(op)
		flag.Int64Var(p, op, value, usage)
	}
}

// UintVar updates c.options doc strings, then splits options and calls flag.Int64Var()
func (c *Cli) UintVar(p *uint, names string, value uint, usage string) {
	// Prep to hand off to the flag package
	ops := splitOps(names)
	// Save for our internal option documentation
	label := opsLabel(ops)
	c.options[label] = usage
	// process with flag package
	for _, op := range ops {
		op = strings.TrimSpace(op)
		flag.UintVar(p, op, value, usage)
	}
}

// Uint64Var updates c.options doc strings, then splits options and calls flag.Int64Var()
func (c *Cli) Uint64Var(p *uint64, names string, value uint64, usage string) {
	// Prep to hand off to the flag package
	ops := splitOps(names)
	// Save for our internal option documentation
	label := opsLabel(ops)
	c.options[label] = usage
	// process with flag package
	for _, op := range ops {
		op = strings.TrimSpace(op)
		flag.Uint64Var(p, op, value, usage)
	}
}

// StringVar updates c.options doc strings, then splits options and calls flag.StringVar()
func (c *Cli) StringVar(p *string, names string, value string, usage string) {
	// Prep to hand off to the flag package
	ops := splitOps(names)
	// Save for our internal option documentation
	label := opsLabel(ops)
	c.options[label] = usage
	// process with flag package
	for _, op := range ops {
		op = strings.TrimSpace(op)
		flag.StringVar(p, op, value, usage)
	}
}

// Float64Var updates c.options doc strings, then splits options and calls flag.Float64Var()
func (c *Cli) Float64Var(p *float64, names string, value float64, usage string) {
	// Prep to hand off to the flag package
	ops := splitOps(names)
	// Save for our internal option documentation
	label := opsLabel(ops)
	c.options[label] = usage
	// process with flag package
	for _, op := range ops {
		op = strings.TrimSpace(op)
		flag.Float64Var(p, op, value, usage)
	}
}

// DurationVar updates c.options doc strings, then splits options and calls flag.DurationVar()
func (c *Cli) DurationVar(p *time.Duration, names string, value time.Duration, usage string) {
	// Prep to hand off to the flag package
	ops := splitOps(names)
	// Save for our internal option documentation
	label := opsLabel(ops)
	c.options[label] = usage
	// process with flag package
	for _, op := range ops {
		op = strings.TrimSpace(op)
		flag.DurationVar(p, op, value, usage)
	}
}

// Option returns an option's document string or unsupported string
func (c *Cli) Option(op string) string {
	op = strings.Trim(op, " ")
	for k, doc := range c.options {
		if strings.Contains(k, op) {
			return doc
		}
	}
	return fmt.Sprintf("%q is an unsupported option", op)
}

// Options returns a map of option values and doc strings
func (c *Cli) Options() map[string]string {
	return c.options
}

// ParseOptions envokes flag.Parse() updating variables set in AddOptions
func (c *Cli) ParseOptions() {
	flag.Parse()
}

// EnvBoolVar adds environment variable which is evaluate before evaluating options
// It is the environment counterpart to flag.BoolVar()
func (c *Cli) EnvBoolVar(p *bool, name string, value bool, usage string) error {
	c.env[name] = &EnvAttribute{
		Name:      name,
		Type:      fmt.Sprintf("%T", value),
		BoolValue: value,
		Usage:     usage,
	}
	*p = c.env[name].BoolValue
	_, ok := c.env[name]
	if ok == false {
		return fmt.Errorf("%q could not be added to environment attributes", name)
	}
	return nil
}

// EnvIntVar adds environment variable which is evaluate before evaluating options
// It is the environment counterpart to flag.IntVar()
func (c *Cli) EnvIntVar(p *int, name string, value int, usage string) error {
	c.env[name] = &EnvAttribute{
		Name:     name,
		Type:     fmt.Sprintf("%T", value),
		IntValue: value,
		Usage:    usage,
	}
	*p = c.env[name].IntValue
	_, ok := c.env[name]
	if ok == false {
		return fmt.Errorf("%q could not be added to environment attributes", name)
	}
	return nil
}

// EnvInt64Var adds environment variable which is evaluate before evaluating options
// It is the environment counterpart to flag.Int64Var()
func (c *Cli) EnvInt64Var(p *int64, name string, value int64, usage string) error {
	c.env[name] = &EnvAttribute{
		Name:       name,
		Type:       fmt.Sprintf("%T", value),
		Int64Value: value,
		Usage:      usage,
	}
	*p = c.env[name].Int64Value
	_, ok := c.env[name]
	if ok == false {
		return fmt.Errorf("%q could not be added to environment attributes", name)
	}
	return nil
}

// EnvUintVar adds environment variable which is evaluate before evaluating options
// It is the environment counterpart to flag.UintVar()
func (c *Cli) EnvUintVar(p *uint, name string, value uint, usage string) error {
	c.env[name] = &EnvAttribute{
		Name:      name,
		Type:      fmt.Sprintf("%T", value),
		UintValue: value,
		Usage:     usage,
	}
	*p = c.env[name].UintValue
	_, ok := c.env[name]
	if ok == false {
		return fmt.Errorf("%q could not be added to environment attributes", name)
	}
	return nil
}

// EnvUint64Var adds environment variable which is evaluate before evaluating options
// It is the environment counterpart to flag.Uint64Var()
func (c *Cli) EnvUint64Var(p *uint64, name string, value uint64, usage string) error {
	c.env[name] = &EnvAttribute{
		Name:        name,
		Type:        fmt.Sprintf("%T", value),
		Uint64Value: value,
		Usage:       usage,
	}
	*p = c.env[name].Uint64Value
	_, ok := c.env[name]
	if ok == false {
		return fmt.Errorf("%q could not be added to environment attributes", name)
	}
	return nil
}

// EnvFloat64Var adds environment variable which is evaluate before evaluating options
// It is the environment counterpart to flag.Float64Var()
func (c *Cli) EnvFloat64Var(p *float64, name string, value float64, usage string) error {
	c.env[name] = &EnvAttribute{
		Name:         name,
		Type:         fmt.Sprintf("%T", value),
		Float64Value: value,
		Usage:        usage,
	}
	*p = c.env[name].Float64Value
	_, ok := c.env[name]
	if ok == false {
		return fmt.Errorf("%q could not be added to environment attributes", name)
	}
	return nil
}

// EnvStringVar adds environment variable which is evaluate before evaluating options
// It is the environment counterpart to flag.StringVar()
func (c *Cli) EnvStringVar(p *string, name string, value string, usage string) error {
	c.env[name] = &EnvAttribute{
		Name:        name,
		Type:        fmt.Sprintf("%T", value),
		StringValue: value,
		Usage:       usage,
	}
	*p = c.env[name].StringValue
	_, ok := c.env[name]
	if ok == false {
		return fmt.Errorf("%q could not be added to environment attributes", name)
	}
	return nil
}

// EnvDurationVar adds environment variable which is evaluate before evaluating options
// It is the environment counterpart to flag.DurationVar()
func (c *Cli) EnvDurationVar(p *time.Duration, name string, value time.Duration, usage string) error {
	c.env[name] = &EnvAttribute{
		Name:          name,
		Type:          fmt.Sprintf("%T", value),
		DurationValue: value,
		Usage:         usage,
	}
	*p = c.env[name].DurationValue
	_, ok := c.env[name]
	if ok == false {
		return fmt.Errorf("%q could not be added to environment attributes", name)
	}
	return nil
}

// EnvAttribute returns the struct corresponding to the matchine name
func (c *Cli) EnvAttribute(name string) (*EnvAttribute, error) {
	e, ok := c.env[name]
	if ok == false {
		return nil, fmt.Errorf("%q not defined for environment", name)
	}
	return e, nil
}

// Env returns an EnvAttribute documentation string for matching name
func (c *Cli) Env(name string) string {
	e, ok := c.env[name]
	if ok == false {
		return fmt.Sprintf("%q not documented for environment", name)
	}
	return e.Usage
}

// Getenv returns a given environment attribute value as a string
func (c *Cli) Getenv(name string) string {
	var s string
	e, err := c.EnvAttribute(name)
	if err != nil {
		return s
	}
	switch e.Type {
	case "bool":
		return fmt.Sprintf("%t", e.BoolValue)
	case "int":
		return fmt.Sprintf("%d", e.IntValue)
	case "int64":
		return fmt.Sprintf("%d", e.Int64Value)
	case "uint":
		return fmt.Sprintf("%d", e.UintValue)
	case "uint64":
		return fmt.Sprintf("%d", e.Uint64Value)
	case "float64":
		return fmt.Sprintf("%f", e.Float64Value)
	case "time.Duration":
		return fmt.Sprintf("%s", e.DurationValue)
	}
	return e.StringValue
}

// ParseEnv loops through the os environment using os.Getenv() and updates
// c.env EnvAttribute. Returns an error if there is a problem with environment.
func (c *Cli) ParseEnv() error {
	var (
		err error
		u64 uint64
	)
	for k, e := range c.env {
		s := strings.TrimSpace(os.Getenv(k))
		switch e.Type {
		case "bool":
			e.BoolValue, err = strconv.ParseBool(s)
		case "int":
			e.IntValue, err = strconv.Atoi(s)
		case "int64":
			e.Int64Value, err = strconv.ParseInt(s, 10, 64)
		case "uint":
			u64, err = strconv.ParseUint(s, 10, 32)
			e.UintValue = uint(u64)
		case "uint64":
			e.Uint64Value, err = strconv.ParseUint(s, 10, 64)
		case "float64":
			e.Float64Value, err = strconv.ParseFloat(s, 64)
		case "time.Duration":
			e.DurationValue, err = time.ParseDuration(s)
		default:
			e.StringValue = s
		}
		if err != nil {
			return fmt.Errorf("%q should be type %q, %s", e.Name, e.Type, err)
		}
		c.env[k] = e
	}
	return err
}

// Parse process both the environment and any flags
func (c *Cli) Parse() error {
	err := c.ParseEnv()
	if err != nil {
		return err
	}
	c.ParseOptions()
	return nil
}

// Args returns flag.Args()
func (c *Cli) Args() []string {
	return flag.Args()
}

// Add Params documents any parameters not defined as Options or Actions, it is an orders list of strings
func (c *Cli) AddParams(params ...string) {
	for _, param := range params {
		c.params = append(c.params, param)
	}
}

// AddAction associates a wrapping function with a action name, the wrapping function
// has 4 parameters in io.Reader, out io.Writer, err io.Writer, args []string. It should return
// an integer reflecting an exit code like you'd pass to os.Exit().
func (c *Cli) AddAction(verb string, fn func(io.Reader, io.Writer, io.Writer, []string) int, usage string) error {
	c.actions[verb] = &Action{
		Name:  verb,
		Fn:    fn,
		Usage: usage,
	}
	_, ok := c.actions[verb]
	if ok == false {
		return fmt.Errorf("Failed to add action %q", verb)
	}
	return nil
}

// Action returns a doc string for a given verb
func (c *Cli) Action(verb string) string {
	action, ok := c.actions[verb]
	if ok == false {
		return fmt.Sprintf("%q is not a defined action", verb)
	}
	return action.Usage
}

// Actions returns a map of actions and their doc strings
func (c *Cli) Actions() map[string]string {
	actions := map[string]string{}
	for k, action := range c.actions {
		actions[k] = action.Usage
	}
	return actions
}

// (c *Cli) Run takes a list of non-option arguments and runs them if the fist arg (i.e. arg[0]
// has a corresponding action. Returns an int suitable to passing to os.Exit()
func (c *Cli) Run(args []string) int {
	if len(args) == 0 {
		fmt.Fprintf(c.Eout, "Nothing to do\n")
		return 1
	}
	verb := strings.TrimSpace(args[0])
	action, ok := c.actions[verb]
	if ok == false {
		fmt.Fprintf(c.Eout, "do not known how to %q\n", verb)
		return 1
	}
	return action.Fn(c.In, c.Out, c.Eout, args)
}

func padRight(s, p string, maxWidth int) string {
	r := []string{s}
	cnt := maxWidth - len(s)
	for i := 0; i < cnt; i++ {
		r = append(r, p)
	}
	return strings.Join(r, "")
}

// Usage writes a help page to io.Writer provided. Documentation is based on
// the application's metadata like app name, version, options, actions, etc.
func (c *Cli) Usage(w io.Writer) {
	var parts []string
	parts = append(parts, c.appName)
	if len(c.options) > 0 {
		parts = append(parts, "[OPTIONS]")
	}
	if len(c.actions) > 0 {
		parts = append(parts, "[ACTION] [ACTION PARAMETERS...]")
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

// GenerateMarkdownDocs writes a Markdown page to io.Writer provided. Documentation is based on
// the application's metadata like app name, version, options, actions, etc.
func (c *Cli) GenerateMarkdownDocs(w io.Writer) {
	var parts []string
	parts = append(parts, c.appName)
	if len(c.options) > 0 {
		parts = append(parts, "[OPTIONS]")
	}
	if len(c.actions) > 0 {
		parts = append(parts, "[ACTION] [ACTION PARAMETERS...]")
	} else if len(c.params) > 0 {
		parts = append(parts, c.params...)
	}
	fmt.Fprintf(w, "\n# USAGE\n\n	%s\n\n", strings.Join(parts, " "))

	if section, ok := c.Documentation["description"]; ok == true {
		fmt.Fprintf(w, "## SYNOPSIS\n\n%s\n\n", section)
	}

	if len(c.env) > 0 {
		fmt.Fprintf(w, "## ENVIRONMENT\n\n")
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
		fmt.Fprintf(w, "```\n")
		for _, k := range keys {
			fmt.Fprintf(w, "    %s  %s\n", padRight(k, " ", padding), c.env[k].Usage)
		}
		fmt.Fprintf(w, "\n\n")
	}

	if len(c.options) > 0 {
		fmt.Fprintf(w, "## OPTIONS\n\n")
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
		fmt.Fprintf(w, "```\n")
		for _, k := range keys {
			fmt.Fprintf(w, "    %s  %s\n", padRight(k, " ", padding), c.options[k])
		}
		fmt.Fprintf(w, "```\n")
		fmt.Fprintf(w, "\n\n")
	}

	if len(c.actions) > 0 {
		fmt.Fprintf(w, "## ACTIONS\n\n")
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
		fmt.Fprintf(w, "```\n")
		for _, k := range keys {
			usage := c.Action(k)
			fmt.Fprintf(w, "    %s  %s\n", padRight(k, " ", padding), usage)
		}
		fmt.Fprintf(w, "```\n")
		fmt.Fprintf(w, "\n\n")
	}

	if section, ok := c.Documentation["examples"]; ok == true {
		fmt.Fprintf(w, "## EXAMPLES\n\n%s\n\n", section)
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

func (c *Cli) License() string {
	license, ok := c.Documentation["license"]
	if ok == true {
		return fmt.Sprintf("%s", license)
	}
	return ""
}

func (c *Cli) Version() string {
	return c.version
}
