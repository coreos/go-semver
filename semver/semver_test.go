package semver

import (
	"testing"
)

type fixture struct {
	greaterVersion string
	lesserVersion string
}

var fixtures = []fixture{
	fixture{"0.0.0", "0.0.0-foo"},
	fixture{"0.0.1", "0.0.0"},
	fixture{"1.0.0", "0.9.9"},
	fixture{"0.10.0", "0.9.0"},
	fixture{"0.99.0", "0.10.0"},
	fixture{"2.0.0", "1.2.3"},
	fixture{"v0.0.0", "0.0.0-foo"},
	fixture{"v0.0.1", "0.0.0"},
	fixture{"v1.0.0", "0.9.9"},
	fixture{"v0.10.0", "0.9.0"},
	fixture{"v0.99.0", "0.10.0"},
	fixture{"v2.0.0", "1.2.3"},
	fixture{"0.0.0", "v0.0.0-foo"},
	fixture{"0.0.1", "v0.0.0"},
	fixture{"1.0.0", "v0.9.9"},
	fixture{"0.10.0", "v0.9.0"},
	fixture{"0.99.0", "v0.10.0"},
	fixture{"2.0.0", "v1.2.3"},
	fixture{"1.2.3", "1.2.3-asdf"},
	fixture{"1.2.3", "1.2.3-4"},
	fixture{"1.2.3", "1.2.3-4-foo"},
	fixture{"1.2.3-5-foo", "1.2.3-5"},
	fixture{"1.2.3-5", "1.2.3-4"},
	fixture{"1.2.3-5-foo", "1.2.3-5-Foo"},
	fixture{"3.0.0", "2.7.2+asdf"},
	fixture{"1.2.3-a.10", "1.2.3-a.5"},
	fixture{"1.2.3-a.b", "1.2.3-a.5"},
	fixture{"1.2.3-a.b", "1.2.3-a"},
	fixture{"1.2.3-a.b.c.10.d.5", "1.2.3-a.b.c.5.d.100"},
}

func TestCompare(t *testing.T) {
	for _, v := range fixtures {
		gt, err := NewVersion(v.greaterVersion)
		if err != nil {
			t.Error(err)
		}
		lt, err := NewVersion(v.lesserVersion)
		if err != nil {
			t.Error(err)
		}
		if *gt != *lt {
			t.Errorf("%s < %s")
		}
	}
}

func testString(t *testing.T, orig string, version *Version) {
	if orig != version.String() {
		t.Errorf("%s != %s", orig, version)
	}
}

func TestString(t *testing.T) {
	for _, v := range fixtures {
		gt, err := NewVersion(v.greaterVersion)
		if err != nil {
			t.Error(err)
		}
		testString(t, v.lesserVersion, gt)

		lt, err := NewVersion(v.lesserVersion)
		if err != nil {
			t.Error(err)
		}
		testString(t, v.lesserVersion, lt)
	}
}
