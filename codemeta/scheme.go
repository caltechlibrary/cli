package codemeta

import (
	"encoding/json"
)

//FIXME: Sketch the fields I need for the outer object,
// See: https://codemeta.github.io/terms/
type CodeMeta struct {
	CodeRepository         *url.URL            `json:"codeRepository,omitempty"`
	ProgrammingLanguage    string              `json:"programmingLanguage,omitempty"`
	RuntimePlatform        string              `json:"runtimePlatform,omitempty"`
	TargetProduct          string              `json:"targetProduct,omitempty"`
	ApplicationCategory    string              `json:"applicationCategory,omitempty"`
	ApplicationSubCategory string              `json:"applicationSubCategory,omitempty"`
	DownloadUrl            *url.URL            `json:"downloadUrl,omitempty"`
	FileSize               string              `json:"fileSize,omitempty"`
	InstallUrl             *url.URL            `json:"installUrl,omitempty"`
	MemoryRequirements     string              `json:"memoryRequirements,omitempty"`
	OperatingSystem        string              `json:"operatingSystem,omitempty"`
	Permissions            string              `json:"permissions,omitempty"`
	ProcessorRequirements  string              `json:"processorRequirements,omitempty"`
	ReleaseNotes           string              `json:"releaseNotes,omitempty"`
	SoftwareHelp           *CreativeWork       `json:"softwareHelp,omitempty"`
	SoftwareRequirements   *SoftwareSourceCode `json:"softwareRequirements,omitempty"`
	SoftwareVersion        string              `json:"softwareVersion,omitempty"`
	SoftwareRequirements   string              `json:"storageRequirements,omitempty"`
	SupportingData         *DataFeed           `json:"supportData,omitempty"`
	Author                 []*Agent            `json:"author,omitempty"`
	Citation               *CreativeWork       `json:"citation,omitempty"`
	Contributor            []*Agent            `json:"contributor,omitempty"`
	CopyrightHolder        *Agent              `json:"copyrightHolder,omitempty"`
	CopyrightYear          int                 `json:"copyrightYear,omitempty"`
	Creator                []*Agent            `json:"creator,omitempty"`
	DateCreated            *time.Time          `json:"dateCreated,omitempty"`
	DateModified           *time.Time          `json:"dateModified,omitempty"`
	DatePublished          *time.Time          `json:"datePublished,omitempty"`
	Editor                 *Person             `json:"editor,omitempty"`
	Encoding               *MediaObject        `json:"encoding,omitempty"`
	FileFormat             string              `json:"fileFormat,omitempty"`
	Funder                 []*Agent            `json:"funder,omitempty"`
	Keywords               []string            `json:"keywords,omitempty"`
	License                string              `json:"license,omitempty"`
	Producer               *Agent              `json:"producer,omitempty"`
	Provider               *Agent              `json:"provider,omitempty"`
	Publisher              *Agent              `json:"publisher,omitempty"`
	Version                *Semvar             `json:"version,omitempty"`
	IsAccessibleForFree    bool                `json:"isAccessibleForFree,omitempty"`
	IsPartOf               *CreativeWork       `json:"isPartOf,omitempty"`
	HasPart                *CreativeWork       `json:"hasPart,omitempty"`
	Position               int                 `json:"position,omitempty"`
	Description            string              `json:"description,omitempty"`
	Identifier             *PropertValue       `json:"identifier,omitempty"`
	Name                   string              `json:"name,omitempty"`
	SameAs                 *url.URL            `json:"sameAs,omitempty"`
	Url                    *url.URL            `json:"url,omitempty"`
	RelatedLink            []*url.URL          `json:"relatedLink,omitempty"`
	GivenName              string              `json:"givenName,omitempty"`
	FamilyName             string              `json:"familyName,omitempty"`
	EMail                  string              `json:"email,omitempty"`
	Affiliation            string              `json:"affiliation,omitempty"`
	Name                   string              `json:"name,omitempty"`
	Address                string              `json:"address,omitempty"`
}
