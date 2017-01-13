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
