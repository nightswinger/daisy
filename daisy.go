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
	return New(time.Now()).BeginningOfDay()
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

func (daisy *Daisy) AsTime() time.Time {
	return daisy.Time
}
