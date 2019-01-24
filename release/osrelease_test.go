package release

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var osreleasetests = []struct {
	in  []string
	out *OSRelease
}{
	{[]string{"NAME=Fedora",
		"VERSION=\"29 (Twenty Nine)\"",
		"ID=fedora",
		"ID_LIKE=linux",
		"VERSION_ID=29",
		"PLATFORM_ID=\"platform:f29\"",
		"PRETTY_NAME=\"Fedora 29 (Twenty Nine)\"",
		"BUILD_ID=build",
		"CPE_NAME=\"cpe:/o:fedoraproject:fedora:29\"",
		"REDHAT_BUGZILLA_PRODUCT=\"Fedora\"",
	},
		&OSRelease{Name: "Fedora", Version: "\"29 (Twenty Nine)\"", ID: "fedora", IDLike: "linux", VersionID: "29", PrettyName: "\"Fedora 29 (Twenty Nine)\"",
		BuildID: "build", CPEName: "\"cpe:/o:fedoraproject:fedora:29\""},
	}}
func TestParseOSRelease(t *testing.T) {
	for _, tt := range osreleasetests {
		t.Run(tt.in[0], func(t *testing.T) {
			testOSRelease := ParseOSRelease(tt.in)
			assert.Equal(t, tt.out, testOSRelease, "structs don't match")
		})
	}
}
