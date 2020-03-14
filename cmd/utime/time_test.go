package utime_test

import (
	"golibs/cmd/utime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseStringToISODate(t *testing.T) {
	dateExpected := time.Date(2018, time.May, 11, 12, 0, 0, 0, time.UTC)

	time, _ := utime.ParseStringToISODate("2018-05-11 12:00:00")

	assert.Equal(t, dateExpected, time)
}

func TestFormatISODateToString(t *testing.T) {
	date := time.Date(2018, time.August, 18, 10, 10, 10, 0, time.UTC)

	valueString := utime.FormatISODateToString(date)

	assert.Equal(t, "2018-08-18 10:10:10", valueString)
}

func TestParseRefundStringToDate(t *testing.T) {
	dateExpected := time.Date(2018, time.March, 12, 18, 56, 16, 0, time.UTC)

	date, _ := utime.ParseStringToRFCDate("2018-03-12T18:56:16Z")

	assert.Equal(t, dateExpected, date)
}

func TestFormatDateToStringWithLocation(t *testing.T) {
	date, _ := utime.ParseStringToRFCDate("2018-03-12T18:56:16Z")

	valueString := utime.FormatDateToStringWithLocation(date)

	assert.Equal(t, "12-MAR-18 15.56.16 PM", valueString)
}

func TestFormatDateToISO8601StringWithLocation(t *testing.T) {
	date, _ := utime.ParseStringToISO8601Date("2018-02-19T14:49:10.377-04:00")

	valueString := utime.FormatDateToISO8601StringWithLocation(date)

	assert.Equal(t, "2018-02-19T15:49:10.377-03:00", valueString)
}

func TestFormatZeroDateToISOStringWithLocation(t *testing.T) {
	valueString := utime.FormatDateToISO8601StringWithLocation(time.Time{})

	assert.Equal(t, "", valueString)
}

func TestFormatDateToFileNameFormat(t *testing.T) {
	date := time.Date(2018, time.March, 20, 18, 56, 16, 0, time.UTC)

	valueString := utime.FormatDateToFileNameFormat(date)

	assert.Equal(t, "2018_03_20_18_56", valueString)
}

func TestFormatErrorNoValidLocation(t *testing.T) {
	date, _ := utime.ParseStringToISO8601Date("2018-02-19T14:49:10.377-04:00")

	valueString := utime.Format(date, "2006-01-02T15:04:05.000-07:00", "invalidLocation")

	assert.Equal(t, "", valueString)
}
