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

var (
	appName = "testcli"
)

func isSameString(expected, value string) string {
	if strings.Compare(expected, value) == 0 {
		return ""
	}
	return fmt.Sprintf("expected %q, got %q", expected, value)
}

func isSameInt(expected, value int) string {
	if value == 0 {
		return ""
	}
	return fmt.Sprintf("expected %d, got %d", expected, value)
}

func check(t *testing.T, msg string, failNow bool) {
	if len(msg) > 0 {
		t.Errorf(msg)
	}
}

func TestNew(t *testing.T) {
	cfg := New(appName, appName, "%s %s released under a BSD License", "v0.0.0")
	check(t, isSameString(appName, cfg.appName), true)
	check(t, isSameString(strings.ToUpper(appName), cfg.EnvPrefix), true)
	check(t, isSameInt(0, len(cfg.Options)), true)
}

func TestOpen(t *testing.T) {
	// Check of use of fallbackFile (os.Stdout) and cli.Open()
	fp, err := Open("", os.Stdout)
	if err != nil {
		t.Errorf("Should have fallen back to os.Stdout, got error %s", err)
		t.FailNow()
	}
	if fp != os.Stdout {
		t.Errorf("fp should be pointing at os.Stdout")
		t.FailNow()
	}
	fp, err = Open("-", os.Stdout)
	if err != nil {
		t.Errorf("Should have fallen back to os.Stdout, got error %s", err)
		t.FailNow()
	}
	if fp != os.Stdout {
		t.Errorf("fp should be pointing at os.Stdout")
		t.FailNow()
	}
	// Check if we can open an existing file, i.e. README.md
	fp, err = Open("README.md", os.Stdout)
	if err != nil {
		t.Errorf("Should have gotten pointer to README.md, got error %s", err)
		t.FailNow()
	}
	if fp == os.Stdout {
		t.Errorf("fp should be pointing at README.md NOT os.Stdout")
		t.FailNow()
	}
	if fp == nil {
		t.Errorf("fp should be pointing at README.md NOT nil")
		t.FailNow()
	}
	fp.Close()

	// Check to see if we fallack to os.Stdout for Create()
	fp, err = Create("", os.Stdout)
	if err != nil {
		t.Errorf("Should have fallen back to os.Stdout, got error %s", err)
		t.FailNow()
	}
	if fp != os.Stdout {
		t.Errorf("fp should be pointing at os.Stdout")
		t.FailNow()
	}
	fp, err = Create("-", os.Stdout)
	if err != nil {
		t.Errorf("Should have fallen back to os.Stdout, got error %s", err)
		t.FailNow()
	}
	if fp != os.Stdout {
		t.Errorf("fp should be pointing at os.Stdout")
		t.FailNow()
	}
	// Check to see if we can open test.txt with Create()
	fp, err = Create("test.txt", os.Stdout)
	if err != nil {
		t.Errorf("Should have gotten pointer to test.txt, got error %s", err)
		t.FailNow()
	}
	if fp == os.Stdout {
		t.Errorf("fp should be pointing at test.txt NOT os.Stdout")
		t.FailNow()
	}
	if fp == nil {
		t.Errorf("fp should be pointing at test.txt NOT nil")
		t.FailNow()
	}
	fp.Close()
	os.Remove("test.txt")
}

func TestMergeEnv(t *testing.T) {
	cfg := New(appName, appName, "", "v0.0.0")

	for _, term := range []string{"API_URL", "DBNAME", "BLEVE", "HTDOCS", "TEMPLATE_PATH", "SITE_URL"} {
		if s := cfg.MergeEnv(term, "test_"+term); strings.Compare(s, "test_"+term) != 0 {
			t.Error(fmt.Sprintf("%s_%s error %s", "EPGO", term, s))
			t.FailNow()
		}
	}
}

func TestMergeBool(t *testing.T) {
	cfg := New("testcli", "TESTCLI", "", "v0.0.0")
	envVar := "TESTCLI_ONOFF"
	onoff := true
	expected := true
	// NOTE: TESTCLI_ONOFF willnoe be set so will always return false
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
