package release

import (
	"strings"
)

// OSRelease contains the contents of /etc/os-release
type OSRelease struct {
	Name             string
	Version          string
	ID               string
	IDLike           string
	VersionID        string
	PrettyName       string
	CPEName          string
	BuildID          string
}

// ParseOSRelease takes a []string of lines in /etc/os-release and returns a filled OSRelease
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
			case "VERSION_ID":
				release.VersionID = split[1]
			case "PRETTY_NAME":
				release.PrettyName = split[1]
			case "CPE_NAME":
				release.CPEName = split[1]
			case "BUILD_ID":
				release.BuildID = split[1]
			default:
			}
		}
	}
	return release
}
