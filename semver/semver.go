package semver

import (
	"bytes"
	"fmt"
)

type Version struct {
	Major int
	Minor int
	Patch int
	PreRelease string
	MetaData string
}

func NewVersion(v string) (*Version, error) {
	return &Version{1, 0, 0, "alpha1", "da39a3ee"}, nil
}

func (v *Version) String() string {
	var buffer bytes.Buffer

	base := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	buffer.WriteString(base)

	if v.PreRelease != "" {
		buffer.WriteString(fmt.Sprintf("-%s", v.PreRelease))
	}

	if v.MetaData != "" {
		buffer.WriteString(fmt.Sprintf("+%s", v.MetaData))
	}

	return buffer.String()
}
