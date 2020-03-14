package utime

import (
	"errors"
	"strings"
	"time"
)

var ErrLocation = errors.New("time: invalid location name")

const (
	FormatISODate = "2006-01-02 15:04:05"
	FileFormat    = "2006_01_02_15_04"
	FormatRFC3339 = "2006-01-02T15:04:05Z"
	FormatISO8601 = "2006-01-02T15:04:05.000-07:00"
)

// ParseStringToISODate parse string to date. Format used is 2006-01-02 15:04:05.
func ParseStringToISODate(date string) (time.Time, error) {
	return time.Parse(FormatISODate, date)
}

// ParseStringToRFCDate parse string to date. Format used is 2006-01-02T15:04:05Z.
func ParseStringToRFCDate(date string) (time.Time, error) {
	return time.Parse(FormatRFC3339, date)
}

// FormatISODateToString format date to string. Format used is 2006-01-02 15:04:05.
func FormatISODateToString(date time.Time) string {
	return date.Format(FormatISODate)
}

// FormatDateToFileNameFormat format date to string. Format used is 2006_01_02_15_04.
func FormatDateToFileNameFormat(date time.Time) string {
	return date.Format(FileFormat)
}

// Format format date to string with specific format and location.
func Format(date time.Time, format string, location string) (dateString string, err error) {
	if !date.IsZero() {
		location, err := time.LoadLocation(location)
		if err != nil {
			if strings.Contains(err.Error(), "unknown time zone") {
				return "", ErrLocation
			}
			return "", err
		}
		dateString = date.UTC().In(location).Format(format)
	}
	return
}
