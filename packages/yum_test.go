package packages

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var yumlinetests = []struct {
	in  string
	out *Package
}{
	{"CharLS.x86_64                                     1.0-17.fc29                            @fedora",
		&Package{Name: "CharLS.x86_64", Version: "1.0-17.fc29"}},
}
func TestParseYumLine(t *testing.T) {
	for _, tt := range yumlinetests {
		t.Run(tt.in, func(t *testing.T) {
			testPackage := ParseYumLine(tt.in)
			assert.Equal(t, tt.out, testPackage, "structs don't match")
		})
	}
}
