package packages

import (
	"strings"
)

// Package is a program installed by a package manager like yum or apt
type Package struct {
	Name    string
	Version string
}

// ParseYumLine returns a Package from a line in `yum list installed`
func ParseYumLine(line string) *Package {
	fields := strings.Fields(line)
	return &Package{Name: fields[0], Version: fields[1]}
}
