package time_utils

import "time"

const apiDateLayout = "2006-01-02T15:04:05Z"

func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString return a string of now
func GetNowString() string {
	return GetNow().UTC().Format(apiDateLayout)
}
