package formattedtime

import (
	"time"
)

// New will return a new instance of FormattedTime
func New(format string) (fp *FormattedTime) {
	var f FormattedTime
	f.format = format
	f.DateTimes = make(map[string]*DateTimeEntry)
	return &f
}

// FormattedTime is the struct of date times to format
type FormattedTime struct {
	format    string
	DateTimes map[string]*DateTimeEntry
}

// NewDateTime will add a DateTimeEntry with an empty DateTime field to the FormattedTime instance
func (f *FormattedTime) NewDateTime(key string) {
	var d DateTimeEntry
	d.formattedTime = f
	f.DateTimes[key] = &d
}

// AddDateTime will add a DateTimeEntry to the FormattedTime instance
func (f *FormattedTime) AddDateTime(key string, dateTime time.Time) {
	var d DateTimeEntry
	d.formattedTime = f

	var dt DateTime
	dt = DateTime(dateTime)
	d.DateTime = &dt

	f.DateTimes[key] = &d
}
