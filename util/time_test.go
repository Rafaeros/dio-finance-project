package util

import "testing"


// TestStringToTime tests that StringToTime returns a time.Time object with the
// correct year, given a string in the format "2006-01-02T15:04:05".
func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2024-10-13T11:01:05")

	if convertedTime.Year() != 2024 {
		t.Errorf("Expected year 2024, got %d", convertedTime.Year())
	}
}