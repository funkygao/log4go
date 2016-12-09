package log4go

import (
	"testing"

	"github.com/funkygao/assert"
)

func TestToLogLevel(t *testing.T) {
	fixtures := assert.Fixtures{
		assert.Fixture{
			Expected: INFO,
			Input:    "info",
		},
		assert.Fixture{
			Expected: TRACE,
			Input:    "trace",
		},
		assert.Fixture{
			Expected: DEBUG,
			Input:    "DebuG",
		},
		assert.Fixture{
			Expected: ERROR,
			Input:    "ERROR",
		},
	}

	for _, f := range fixtures {
		lvl := ToLogLevel(f.Input.(string), CRITICAL)
		assert.Equal(t, f.Expected.(Level), lvl)
	}
}
