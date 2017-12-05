/**
 * cli is a package intended to encourage some standardization in the command line user interface for programs
 * developed for Caltech Library.
 *
 * @author R. S. Doiel, <rsdoiel@caltech.edu>
 *
 * Copyright (c) 2017, Caltech
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
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	licenseString := "%s %s released under a BSD License"
	cfg := New(appName, strings.ToUpper(appName), "v0.0.0")
	cfg.LicenseText = strings.Replace(licenseString, "%s", appName, -1)
	check(t, isSameString(appName, cfg.appName), true)
	check(t, isSameString(strings.ToUpper(appName), cfg.EnvPrefix), true)
	check(t, isSameInt(0, len(cfg.Options)), true)
}

func TestMergeEnv(t *testing.T) {
	cfg := New(appName, appName, "v0.0.0")

	for _, term := range []string{"API_URL", "DBNAME", "BLEVE", "HTDOCS", "TEMPLATE_PATH", "SITE_URL"} {
		if s := cfg.MergeEnv(term, "test_"+term); strings.Compare(s, "test_"+term) != 0 {
			t.Error(fmt.Sprintf("%s_%s error %s", "EPGO", term, s))
			t.FailNow()
		}
	}
}

func TestMergeBool(t *testing.T) {
	cfg := New("testcli", "TESTCLI", "v0.0.0")
	envVar := "TESTCLI_ONOFF"
	onoff := true
	expected := true
	// NOTE: TESTCLI_ONOFF will not be set so will always return false
	result := cfg.MergeEnvBool("onoff", onoff)
	if expected != result {
		t.Errorf("For onoff %t (env: %q), Expected %t, got %t", onoff, os.Getenv(envVar), expected, result)
	}

	onoff = false
	expected = false
	result = cfg.MergeEnvBool("onoff", onoff)
	if expected != result {
		t.Errorf("For onoff %t (env: %q), Expected %t, got %t", onoff, os.Getenv(envVar), expected, result)
	}
}

func TestStandardOptions(t *testing.T) {
	cfg := New("testcli", "TESTCLI", "v0.0.0")
	args := []string{}
	showHelp := false
	showExamples := false
	showLicense := false
	showVersion := false
	text := cfg.StandardOptions(showHelp, showExamples, showLicense, showVersion, args)
	if text != "" {
		t.Errorf("Expected empty string, --->\n%s\n", text)
		t.FailNow()
	}
	showHelp = true
	showExamples = false
	args = append(args, "append")
	text = cfg.StandardOptions(showHelp, showExamples, showLicense, showVersion, args)
	if strings.HasPrefix(text, "No Information for append") {
		t.Errorf("Expected 'No information for append', got %q", text)
		t.FailNow()
	}
}
