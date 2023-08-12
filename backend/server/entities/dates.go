package entities

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/imartinezalberte/go-todo-list/backend/utils"
)

const (
	// Date format in human representation
	HumanDateFormat     = "2006-01-02"
	HumanHourFormat     = "15:04"
	FullHumanDateFormat = HumanDateFormat + utils.Space + HumanHourFormat
)

type (
	DateYearMonthDayHourMinute time.Time

	DateYearMonthDay time.Time

	DateHourMinute time.Time
)

func (d *DateYearMonthDayHourMinute) UnmarshalJSON(input []byte) error {
	t, err := UnmarshalDateJSON(input, FullHumanDateFormat)
	*d = DateYearMonthDayHourMinute(t)
	return err
}

func (d DateYearMonthDayHourMinute) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(t.Format(FullHumanDateFormat))
}

func (d DateYearMonthDayHourMinute) String() string {
	return time.Time(d).Format(FullHumanDateFormat)
}

func (d *DateYearMonthDay) UnmarshalJSON(input []byte) error {
	t, err := UnmarshalDateJSON(input, HumanDateFormat)
	*d = DateYearMonthDay(t)
	return err
}

func (d DateYearMonthDay) MarshalJSON() ([]byte, error) {
	return MarshalDateJSON(time.Time(d), HumanDateFormat)
}

func (d DateYearMonthDay) String() string {
	return time.Time(d).Format(HumanDateFormat)
}

func (d *DateHourMinute) UnmarshalJSON(input []byte) error {
	t, err := UnmarshalDateJSON(input, HumanHourFormat)
	*d = DateHourMinute(t)
	return err
}

func (d DateHourMinute) MarshalJSON() ([]byte, error) {
	return MarshalDateJSON(time.Time(d), HumanHourFormat)
}

func (d DateHourMinute) String() string {
	return time.Time(d).Format(HumanHourFormat)
}

// helpers
func UnmarshalDateJSON(input []byte, layout string) (time.Time, error) {
	sanitized := strings.Trim(string(input), utils.Quote)
	if strings.TrimSpace(sanitized) == "" {
		return time.Time{}, nil
	}

	return time.Parse(layout, sanitized)
}

func MarshalDateJSON(input time.Time, layout string) ([]byte, error) {
	if input.IsZero() {
		return json.Marshal(nil)
	}
	return json.Marshal(input.Format(layout))
}
