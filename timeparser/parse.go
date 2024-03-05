package timeparser

import (
	"errors"
	"time"
)

var (
	ErrTimeNotParsed = errors.New("unable to parse time")
)

const (
	// Replicated time format
	//
	// The time format used by Replicated.
	Replicated = "2006-01-02 15:04:05.999999999 -0700 MST"
)

//nolint:gochecknoglobals // slice and struct constants.
var (
	timeFormats = []string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.Kitchen,
		time.Layout,
		// Handy time stamps.
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		time.DateTime,
		Replicated,
	}

	timeZero = time.Time{}
)

func Parse(in string) (time.Time, error) {
	for _, tf := range timeFormats {
		ts, err := time.Parse(tf, in)
		if err == nil {
			return ts, nil
		}
	}
	return timeZero, ErrTimeNotParsed
}
