package user

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var passwdlinetests = []struct {
	in  string
	out *User
}{
	{"nginx:x:990:978:Nginx web server:/var/lib/nginx:/sbin/nologin",
		&User{Username: "nginx", Password: "x", UserID: 990, GroupID: 978, Comment: "Nginx web server", HomeDirectory: "/var/lib/nginx",
			Shell: "/sbin/nologin"}},
}
func TestParsePasswdLine(t *testing.T) {
	for _, tt := range passwdlinetests {
		t.Run(tt.in, func(t *testing.T) {
			testUser := ParsePasswdLine(tt.in)
			assert.Equal(t, tt.out, testUser, "structs don't match")
		})
	}
}
