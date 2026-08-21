package timestamp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConverToHour(t *testing.T) {
	ts := 1786480950601377300
	require.Equal(t, "2:12", ToHourMinute(int64(ts)))
}

func TestConvertMinutes(t *testing.T) {
	ts := 1786570773049333246
	t.Log(ToHourMinute(int64(ts)))
}
