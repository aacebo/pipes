package utils

import (
	"fmt"
	"math"
	"time"
)

func FormatTimeFromNow(t *time.Time) string {
	diff := time.Now().Sub(*t)
	hours := math.Round(diff.Hours())
	mins := math.Round(diff.Minutes())
	secs := math.Round(diff.Seconds())
	days := math.Round(hours / 24)
	months := math.Round(days / 30)

	if months >= 1 {
		return fmt.Sprintf("%.0f months", months)
	} else if days >= 1 {
		return fmt.Sprintf("%.0f days", days)
	} else if hours >= 1 {
		return fmt.Sprintf("%.0f hours", hours)
	} else if mins >= 1 {
		return fmt.Sprintf("%.0f minutes", mins)
	}

	return fmt.Sprintf("%.0f seconds", secs)
}
