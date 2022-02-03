// Package convertdhms provides a converter that process minutes to dd:hh:mm:ss format string
package converttodhms

import (
	"fmt"
	"math"
)

// getMinutesToDhms receive a minutes in integer or float representation and convert it to dd:hh:mm:ss
// (days, hours, minutes, seconds) format string
// Ex: 9623 = 06:16:22:59
func GetMinutesToDhms(input float64) string {
	///home/nanan/convert-to-dhms
	var days, hours, mins, secs float64
	strTime := "00:00:00:00"

	if input > 0 {
		// Day part - start
		days = input / float64(1440)
		stPart := math.Floor(days)
		sdPart := math.Abs(stPart - days)
		days = stPart
		// Day part - end

		// Hour part - start
		hours = sdPart * 24
		stPart = math.Floor(hours)
		sdPart = math.Abs(stPart - hours)
		hours = stPart
		// Hour part - end

		// Minute part - start
		mins = sdPart * 60
		stPart = math.Floor(mins)
		sdPart = math.Abs(stPart - mins)
		mins = stPart
		// Minute part - end

		// Second part - start
		secs = sdPart * 60
		stPart = math.Floor(secs)
		secs = stPart
		// Second part - end

		return fmt.Sprintf("%v:%v:%v:%v", pad(days), pad(hours), pad(mins), pad(secs))

	}

	return strTime
}

// Check if a number is less than 10 and set a zero to the left returning a string with a integer value representation
// If is greater than 10 or equals just put the number like an integer without zero
func pad(in float64) string {
	if in < 10 {
		return fmt.Sprintf("0%.f", in)
	}

	return fmt.Sprintf("%.f", in)
}
