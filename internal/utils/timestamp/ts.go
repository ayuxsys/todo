package timestamp

import (
	"fmt"
	"time"
)

func ToHourMinute(ts int64) string {
	t := time.Unix(0, ts)
	return fmt.Sprintf("%d:%02d", t.Hour(), t.Minute())
}
