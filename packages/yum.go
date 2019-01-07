package packages

import (
	"strings"
)

type Package struct {
	Name    string
	Version string
}

// ParseYumLine returns a Package from a line in `yum list installed`
func ParseYumLine(line string) *Package {
	fields := strings.Fields(line)
	return &Package{Name: fields[0], Version: fields[1]}
}
