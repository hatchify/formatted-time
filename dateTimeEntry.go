package formattedtime

import "time"

// DateTimeEntry represents a date time entry to be formatted
type DateTimeEntry struct {
	formattedTime *FormattedTime
	DateTime      *DateTime
}

// MarshalJSON is a helper method for marshalling JSON
func (d DateTimeEntry) MarshalJSON() (bs []byte, err error) {
	t := d.DateTime.Time()
	str := "\"" + t.Format(d.formattedTime.format) + "\""
	bs = []byte(str)
	return
}

// UnmarshalJSON is a helper method for unmarshalling JSON
func (d *DateTimeEntry) UnmarshalJSON(bs []byte) (err error) {
	if ok := len(bs) > 2; !ok {
		return
	}

	bs = bs[1 : len(bs)-1]

	var t time.Time
	if t, err = time.Parse(d.formattedTime.format, string(bs)); err != nil {
		return
	}

	var dt DateTime
	dt = DateTime(t)

	d.DateTime = &dt
	return
}
