package daisy

import "time"

var WeekStartDay = time.Monday

type Daisy struct {
	time.Time
}

func New(t time.Time) *Daisy {
	return &Daisy{t}
}

func Now() *Daisy {
	return New(time.Now())
}

func Unix(sec int64) *Daisy {
	return New(time.Unix(sec, 0))
}

func Parse(layout, value string) (*Daisy, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}
	return New(t), err
}

func MustParse(layout, value string) *Daisy {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return New(t)
}

func (daisy *Daisy) AddHours(i int) *Daisy {
	return New(daisy.Time.Add(time.Duration(i) * time.Hour))
}

func (daisy *Daisy) AddDays(i int) *Daisy {
	return New(daisy.Time.AddDate(0, 0, i))
}

func (daisy *Daisy) Tomorrow() *Daisy {
	return New(daisy.Time.AddDate(0, 0, 1))
}

func (daisy *Daisy) Yesterday() *Daisy {
	return New(daisy.Time.AddDate(0, 0, -1))
}

func (daisy *Daisy) PrevYear() *Daisy {
	return New(daisy.Time.AddDate(-1, 0, 0))
}

func (daisy *Daisy) NextYear() *Daisy {
	return New(daisy.Time.AddDate(1, 0, 0))
}

func (daisy *Daisy) PrevMonth() *Daisy {
	return New(daisy.Time.AddDate(0, -1, 0))
}

func (daisy *Daisy) NextMonth() *Daisy {
	return New(daisy.Time.AddDate(0, 1, 0))
}

func (daisy *Daisy) PrevWeek() *Daisy {
	return New(daisy.Time.AddDate(0, 0, -7))
}

func (daisy *Daisy) NextWeek() *Daisy {
	return New(daisy.Time.AddDate(0, 0, 7))
}

func (daisy *Daisy) Unix() int64 {
	return daisy.Time.Unix()
}

func (daisy *Daisy) RFC3339() string {
	return daisy.Time.Format(time.RFC3339)
}

func (daisy *Daisy) BeginningOfHour() *Daisy {
	y, m, d := daisy.Date()
	return New(time.Date(y, m, d, daisy.Time.Hour(), 0, 0, 0, daisy.Time.Location()))
}

func (daisy *Daisy) BeginningOfDay() *Daisy {
	y, m, d := daisy.Date()
	return New(time.Date(y, m, d, 0, 0, 0, 0, daisy.Time.Location()))
}

func (daisy *Daisy) BeginningOfWeek() *Daisy {
	weekday := int(daisy.Time.Weekday())

	if WeekStartDay != time.Sunday {
		weekStartDayInt := int(WeekStartDay)

		if weekday < weekStartDayInt {
			weekday += 7 - weekStartDayInt
		} else {
			weekday -= weekStartDayInt
		}
	}
	return New(daisy.Time.AddDate(0, 0, -weekday)).BeginningOfDay()
}

func (daisy *Daisy) BeginningOfMonth() *Daisy {
	y, m, _ := daisy.Date()
	return New(time.Date(y, m, 1, 0, 0, 0, 0, daisy.Location()))
}

func (daisy *Daisy) BeginningOfYear() *Daisy {
	y, _, _ := daisy.Date()
	return New(time.Date(y, 1, 1, 0, 0, 0, 0, daisy.Location()))
}

func (daisy *Daisy) EndOfHour() *Daisy {
	return New(daisy.BeginningOfHour().Time.Add(time.Hour - time.Nanosecond))
}

func (daisy *Daisy) EndOfDay() *Daisy {
	y, m, d := daisy.Date()
	return New(time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), daisy.Location()))
}

func (daisy *Daisy) EndOfWeek() *Daisy {
	return New(daisy.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond))
}

func (daisy *Daisy) EndOfMonth() *Daisy {
	return New(daisy.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond))
}

func BeginningOfHour() *Daisy {
	return New(time.Now()).BeginningOfHour()
}

func BeginningOfDay() *Daisy {
	return New(time.Now()).BeginningOfDay()
}

func BeginningOfWeek() *Daisy {
	return New(time.Now()).BeginningOfWeek()
}

func BeginningOfMonth() *Daisy {
	return New(time.Now()).BeginningOfMonth()
}

func BeginningOfYear() *Daisy {
	return New(time.Now()).BeginningOfYear()
}

func (daisy *Daisy) AsTime() time.Time {
	return daisy.Time
}
