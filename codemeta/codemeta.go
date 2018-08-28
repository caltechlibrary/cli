package codemeta

import (

	// Caltech Library packages
	"github.com/caltechlibrary/cli"
)

const (
	Version = cli.Version

	LicenseText = cli.LicenseText
)

var (
	// Aliases implements common crosswalks to dotpath
	// in the codemeta.json file. This intended to make it
	// easier for those who are learning the codemeta.json
	// standards.
	Aliases = map[string]string{
		"@type":               ".@type",
		"author.given_name":   ".authors[0].given_name",
		"author.family_name":  ".authors[0].family_name",
		"author.email":        ".author[0].email",
		"author.orcid_url":    ".author[0].@id",
		"who":                 ".author[0].name",
		"full_name":           ".author[0].name",
		"readme":              ".readme",
		"what":                ".description",
		"description":         ".description",
		"version":             ".version",
		"repo":                ".codeRepository",
		"repository":          ".codeRepository",
		"codeRepository":      ".codeRepository",
		"issues":              ".issueTracker",
		"issueTracker":        ".issueTracker",
		"license":             ".license",
		"status":              ".developmentStatus",
		"developmentStatus":   ".developmentStatus",
		"releases":            ".downloadUrl",
		"downloadUrl":         ".downloadUrl",
		"keywords":            ".keywords",
		"orcid_url":           ".maintainer",
		"maintainer":          ".maintainer",
		"programmingLanguage": ".programmingLanguage",
	}
)
