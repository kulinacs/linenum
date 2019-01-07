package user

import (
	"testing"
)

func TestParsePasswdLine(t *testing.T) {
	testData := "nginx:x:990:978:Nginx web server:/var/lib/nginx:/sbin/nologin"
	testUser := ParsePasswdLine(testData)
	if testUser.Username != "nginx" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", testUser.Username, "nginx")
	}
	if testUser.Password != "x" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", testUser.Password, "x")
	}
	if testUser.UserID != 990 {
		t.Errorf("Received value was incorrect, got: %d, want: %d", testUser.UserID, 990)
	}
	if testUser.GroupID != 978 {
		t.Errorf("Received value was incorrect, got: %d, want: %d", testUser.GroupID, 978)
	}
	if testUser.Comment != "Nginx web server" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", testUser.Comment, "Nginx web server")
	}
	if testUser.HomeDirectory != "/var/lib/nginx" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", testUser.HomeDirectory, "/var/lib/nginx")
	}
	if testUser.Shell != "/sbin/nologin" {
		t.Errorf("Received value was incorrect, got: %s, want: %s", testUser.Shell, "/sbin/nologin")
	}
}
