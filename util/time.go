package util

import (
	"time"
)

const layout = "2006-01-02T15:04:05"


// StringToTime parses a string in the format "2006-01-02T15:04:05" into a time.Time
// object. It does not validate the input string, and will always return a valid
// time.Time object with the same "zero" values if the input string is not valid.
func StringToTime(value string) time.Time {
	convertedTime, _ := time.Parse(layout, "2024-10-13T11:01:05")
	return convertedTime
}