package utime

import (
	"strings"
	"time"
)

const (
	formatISODate  = "2006-01-02 15:04:05"
	stringLocation = "America/Argentina/Buenos_Aires"
	fileFormat     = "2006_01_02_15_04"
	formatRFC3339  = "2006-01-02T15:04:05Z"
	formatISO8601  = "2006-01-02T15:04:05.000-07:00"
	formatAMPMDate = "02-Jan-06 15.04.05 PM"
)

// ParseStringToISODate parse string to date. Format used is 2006-01-02 15:04:05.
func ParseStringToISODate(date string) (time.Time, error) {
	return time.Parse(formatISODate, date)
}

// ParseStringToISO8601Date parse string to date. Format used is 2006-01-02T15:04:05.000-07:00.
func ParseStringToISO8601Date(date string) (time.Time, error) {
	return time.Parse(formatISO8601, date)
}

// ParseStringToRFCDate parse string to date. Format used is 2006-01-02T15:04:05Z.
func ParseStringToRFCDate(date string) (time.Time, error) {
	return time.Parse(formatRFC3339, date)
}

// FormatISODateToString format date to string. Format used is 2006-01-02 15:04:05.
func FormatISODateToString(date time.Time) string {
	return date.Format(formatISODate)
}

// FormatDateToStringWithLocation format date to string with location (Buenos Aires). Format used is 02-Jan-06 15.04.05 PM.
func FormatDateToStringWithLocation(date time.Time) string {
	return strings.ToUpper(Format(date, formatAMPMDate, stringLocation))
}

// FormatDateToISO8601StringWithLocation format date to string with location (Buenos Aires). Format used is 2006-01-02 15:04:05.
func FormatDateToISO8601StringWithLocation(date time.Time) string {
	return Format(date, formatISO8601, stringLocation)
}

// FormatDateToFileNameFormat format date to string. Format used is 2006_01_02_15_04.
func FormatDateToFileNameFormat(date time.Time) string {
	return date.Format(fileFormat)
}

// Format format date to string with specific format and location.
func Format(date time.Time, format string, location string) (dateString string) {
	if !date.IsZero() {
		location, err := time.LoadLocation(location)
		if err != nil {
			return ""
		}
		dateString = date.UTC().In(location).Format(format)
	}
	return
}
