package util

import (
	"fmt"
	"time"
)

func FormatDate(d time.Time) string {
	return fmt.Sprintf(
		"%d-%02d-%02d %02d:%02d:%02d",
		d.Year(), d.Month(), d.Day(),
		d.Hour(), d.Minute(), d.Second(),
	)
}
