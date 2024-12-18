package util

import (
	"fmt"
	"time"
)

func GetRelativeDuration(date time.Time) (string, error) {
	today := time.Now()
	if date.After(today) {
		return "", fmt.Errorf("GetRelativePastTime: got date before today: %s", date.String())
	}

	duration := today.Sub(date)
	days := int(duration.Hours() / 24)

	result := ""
	if days < 0 {
		panic("difference between today and previous day should never be less than 0")
	}

	// eh no one cares if we're a day or two off right??
	if days <= 6 {
		result = fmt.Sprintf("Uploaded %d Days Ago", days)
	} else if days <= 30 {
		result = fmt.Sprintf("Uploaded %d Weeks Ago", days/7)
	} else if days <= 365 {
		result = fmt.Sprintf("Uploaded %d Months Ago", days/30)
	} else {
		result = fmt.Sprintf("Uploaded %d Years Ago", days/365)
	}

	return result, nil
}

func GetReadingTime(wordCount int) int {
	// not handling hours or seconds because I'm unlikely to ever write something this long
	// 200 is slightly below the average reading time (source: google)
	return wordCount / 200
}
