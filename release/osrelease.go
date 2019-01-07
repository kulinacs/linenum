package release

import (
	"strings"
)

type OSRelease struct {
	Name             string
	Version          string
	ID               string
	IDLike           string
	VersionCodename  string
	VersionID        string
	PrettyName       string
	ANSIColor        string
	CPEName          string
	HomeURL          string
	DocumentationURL string
	SupportURL       string
	BugReportURL     string
	PrivacyPolicyURL string
	BuildID          string
	Variant          string
	VariantID        string
	Logo             string
}

func ParseOSRelease(releaseLines []string) *OSRelease {
	release := &OSRelease{}
	for _, line := range releaseLines {
		split := strings.SplitN(line, "=", 2)
		if len(split) == 2 {
			switch split[0] {
			case "NAME":
				release.Name = split[1]
			case "VERSION":
				release.Version = split[1]
			case "ID":
				release.ID = split[1]
			case "ID_LIKE":
				release.IDLike = split[1]
			case "VERSION_CODENAME":
				release.VersionCodename = split[1]
			case "VERSION_ID":
				release.VersionID = split[1]
			case "PRETTY_NAME":
				release.PrettyName = split[1]
			case "ANSI_COLOR":
				release.ANSIColor = split[1]
			case "CPE_NAME":
				release.CPEName = split[1]
			case "HOME_URL":
				release.HomeURL = split[1]
			case "DOCUMENTATION_URL":
				release.DocumentationURL = split[1]
			case "SUPPORT_URL":
				release.SupportURL = split[1]
			case "BUG_REPORT_URL":
				release.BugReportURL = split[1]
			case "PRIVACY_POLICY_URL":
				release.PrivacyPolicyURL = split[1]
			case "BUILD_ID":
				release.BuildID = split[1]
			case "VARIANT":
				release.Variant = split[1]
			case "VARIANT_ID":
				release.VariantID = split[1]
			case "LOGO":
				release.Logo = split[1]
			default:
			}
		}
	}
	return release
}
