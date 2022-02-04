// Package convertdhms provides a converter that process minutes to dd:hh:mm:ss format string
package converttodhms

import (
	"testing"
)

func TestGetMinutesToDhms(t *testing.T) {
	v := GetMinutesToDhms(9623)
	if v != "06:16:22:59" {
		t.Error("Expected", "06:16:22:59", "Got", v)
	}

	v = GetMinutesToDhms(1440)
	if v != "01:00:00:00" {
		t.Error("Expected", "01:00:00:00", "Got", v)
	}

	v = GetMinutesToDhms(40)
	if v != "00:00:40:00" {
		t.Error("Expected", "00:00:40:00", "Got", v)
	}

	v = GetMinutesToDhms(641)
	if v != "00:10:41:00" {
		t.Error("Expected", "00:10:41:00", "Got", v)
	}
}
