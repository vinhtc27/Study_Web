package log

import (
	"time"
)

func GetFilenameDate() string {
	// Use layout string for time format.
	const layout = "01-02-2006"
	// Place now in the string.
	t := time.Now()
	return "file-" + t.Format(layout) + ".txt"
}
