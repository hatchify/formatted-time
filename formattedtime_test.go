package formattedtime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_New(t *testing.T) {
	var f *FormattedTime

	format := "2006-01-02 03:04 PM"
	f = New(format)
	fmt.Println(f)
}

func Test_AddDateTimeEntry(t *testing.T) {
	var f *FormattedTime

	format := "2006-01-02 03:04 PM"
	f = New(format)

	now := time.Now()

	key := "Test"
	dateTime := now
	f.AddDateTime(key, dateTime)

	dt := f.DateTimes["Test"]
	if time.Time(dt.DateTime.Time()) != now {
		t.Fatalf("times do not match: expected dt to be equal to now as %+v, but received %+v", now, time.Time(dt.DateTime.Time()))
	}
}

func Test_MarshalJSON(t *testing.T) {
	var (
		f   *FormattedTime
		err error
	)

	format := "2006-01-02 03:04 PM"
	f = New(format)

	payload := testNewMarshal(f)

	var bodyBytes []byte
	if bodyBytes, err = json.Marshal(payload); err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(bodyBytes))
}

func Test_UnmarshalJSON(t *testing.T) {
	var (
		f   *FormattedTime
		err error
	)

	format := "2006-01-02 03:04 PM"
	f = New(format)

	payload := testNewMarshal(f)

	var bodyBytes []byte
	if bodyBytes, err = json.Marshal(payload); err != nil {
		t.Fatal(err)
	}

	test := testNewUnmarshal()

	if err = json.Unmarshal(bodyBytes, &test); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("test: %+v", test.DueDate.DateTime.Time())
}

func testNewMarshal(f *FormattedTime) (request TestMarshal) {
	f.AddDateTime("dueDate", time.Now())
	request.DueDate = f.DateTimes["dueDate"]
	return
}

type TestMarshal struct {
	DueDate *DateTimeEntry `json:"dueDate"`
}

func testNewUnmarshal() (request TestUnmarshal) {
	var f *FormattedTime

	format := "2006-01-02 03:04 PM"
	f = New(format)

	key := "dueDate"
	f.NewDateTime(key)

	request.DueDate = f.DateTimes[key]
	return
}

type TestUnmarshal struct {
	DueDate *DateTimeEntry `json:"dueDate"`
}
