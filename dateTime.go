package formattedtime

import "time"

// DateTime is the helper type for parsing date times
type DateTime time.Time

// Time returns time.Time
func (d *DateTime) Time() time.Time {
	return time.Time(*d)
}
