package issues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxGroupSeverity(t *testing.T) {
	ii := Outcome{
		"s1": Issues{
			New(Root, OkLevel, "i1"),
		},
		"s2": Issues{
			New(Root, OkLevel, "i1"),
			New(Root, WarnLevel, "i2"),
			New("g1", WarnLevel, "i2"),
		},
	}

	assert.Equal(t, OkLevel, ii.MaxGroupSeverity("s1", Root))
	assert.Equal(t, WarnLevel, ii.MaxGroupSeverity("s2", Root))
}

func TestIssuesForGroup(t *testing.T) {
	ii := Outcome{
		"s1": Issues{
			New(Root, OkLevel, "i1"),
		},
		"s2": Issues{
			New(Root, OkLevel, "i1"),
			New(Root, WarnLevel, "i2"),
			New("g1", WarnLevel, "i3"),
			New("g1", WarnLevel, "i4"),
		},
	}

	assert.Equal(t, 1, len(ii.For("s1", Root)))
	assert.Equal(t, 2, len(ii.For("s2", "g1")))
}
