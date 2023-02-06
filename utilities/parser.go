package utilities

import (
	"fmt"
	"time"
)

func ParseUnixTime(unixTime uint64) string {
  	t := time.Unix(int64(unixTime) / 1000, 0)
    return t.Format(time.RFC822) 
}

func FmtDuration(seconds uint64) string {
	timeString := fmt.Sprintf("%ds", seconds)
	d, _ := time.ParseDuration(timeString)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	if h == 0 {
		return fmt.Sprintf("%02d:%02d", m, s)
	} else {
		return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	}
}

func HumanReadable(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %c", float64(b)/float64(div), "kMBTQ"[exp])
}
