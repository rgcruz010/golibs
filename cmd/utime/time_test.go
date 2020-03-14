package utime_test

import (
	"golibs/cmd/utime"
	"reflect"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	type args struct {
		date     time.Time
		format   string
		location string
	}
	tests := []struct {
		name           string
		args           args
		wantDateString string
		wantErr        bool
	}{
		{
			name: "Format ISO8601 for America/Argentina/Buenos_Aires",
			args: args{
				date:     time.Date(2018, time.March, 20, 15, 49, 10, 0, time.UTC),
				format:   utime.FormatISO8601,
				location: "America/Argentina/Buenos_Aires",
			},
			wantDateString: "2018-03-20T12:49:10.000-03:00",
			wantErr:        false,
		},
		{
			name: "Format RFC3339 for America/Argentina/Buenos_Aires",
			args: args{
				date:     time.Date(2018, time.March, 20, 15, 49, 10, 0, time.UTC),
				format:   utime.FormatRFC3339,
				location: "America/Argentina/Buenos_Aires",
			},
			wantDateString: "2018-03-20T12:49:10Z",
			wantErr:        false,
		},
		{
			name: "Format FileFormat for America/Argentina/Buenos_Aires",
			args: args{
				date:     time.Date(2018, time.March, 20, 15, 49, 10, 0, time.UTC),
				format:   utime.FileFormat,
				location: "America/Argentina/Buenos_Aires",
			},
			wantDateString: "2018_03_20_12_49",
			wantErr:        false,
		},
		{
			name: "Format FileFormat for invalid Location America/Argentina/Buenos_Air",
			args: args{
				date:     time.Date(2018, time.March, 20, 15, 49, 10, 0, time.UTC),
				format:   utime.FileFormat,
				location: "America/Argentina/Buenos_Air",
			},
			wantDateString: "",
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDateString, err := utime.Format(tt.args.date, tt.args.format, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDateString != tt.wantDateString {
				t.Errorf("Format() gotDateString = %v, want %v", gotDateString, tt.wantDateString)
			}
		})
	}
}

func TestFormatISODateToString(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				date: time.Date(2018, time.August, 18, 10, 10, 10, 0, time.UTC),
			},
			want: "2018-08-18 10:10:10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utime.FormatISODateToString(tt.args.date); got != tt.want {
				t.Errorf("FormatISODateToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseStringToISODate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "",
			args: args{
				date: "2018-05-11 12:00:00",
			},
			want:    time.Date(2018, time.May, 11, 12, 0, 0, 0, time.UTC),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utime.ParseStringToISODate(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStringToISODate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStringToISODate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseStringToRFCDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "",
			args: args{
				date: "2018-03-12T18:56:16Z",
			},
			want:    time.Date(2018, time.March, 12, 18, 56, 16, 0, time.UTC),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utime.ParseStringToRFCDate(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStringToRFCDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStringToRFCDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatDateToFileNameFormat(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				date: time.Date(2018, time.March, 20, 18, 56, 16, 0, time.UTC),
			},
			want: "2018_03_20_18_56",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utime.FormatDateToFileNameFormat(tt.args.date)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStringToRFCDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
