package user

import (
	"strconv"
	"strings"
)

// User is a user on the system found in /etc/passwd
type User struct {
	Username      string
	Password      string
	UserID        int
	GroupID       int
	Groups        []int
	Comment       string
	HomeDirectory string
	Shell         string
}

// ParsePasswdLine returns a User from a line in /etc/passwd
func ParsePasswdLine(line string) *User {
	split := strings.Split(line, ":")
	uid, _ := strconv.Atoi(split[2])
	gid, _ := strconv.Atoi(split[3])
	return &User{Username: split[0], Password: split[1], UserID: uid, GroupID: gid,
		Comment: split[4], HomeDirectory: split[5], Shell: split[6]}
}
