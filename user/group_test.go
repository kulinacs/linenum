package user

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var grouplinetests = []struct {
	in  string
	out *Group
}{
	{"nginx:x:978:", &Group{Groupname: "nginx", Password: "x", GroupID: 978}},
}
func TestParseGroupLine(t *testing.T) {
	for _, tt := range grouplinetests {
		t.Run(tt.in, func(t *testing.T) {
			testGroup := ParseGroupLine(tt.in)
			assert.Equal(t, tt.out, testGroup, "structs don't match")
		})
	}
}
