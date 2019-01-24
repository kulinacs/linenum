package user

import (
	"strings"
	"strconv"
)

// Group is a group on the system found in /etc/group
type Group struct {
	Groupname string
	Password  string
	GroupID   int
}

// ParseGroupLine returns a Group from a line in /etc/group
func ParseGroupLine(line string) *Group {
	split := strings.Split(line, ":")
	gid, _ := strconv.Atoi(split[2])
	return &Group{Groupname: split[0], Password: split[1], GroupID: gid,}
}
