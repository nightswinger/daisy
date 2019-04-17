package daisy

import "time"

type Daisy struct {
	time.Time
}

func New(t time.Time) *Daisy {
	return &Daisy{t}
}

func Now() *Daisy {
	return New(time.Now()).BeginningOfDay()
}

func (daisy *Daisy) Tommorrow() *Daisy {
	return New(daisy.Time.AddDate(0, 0, 1)).BeginningOfDay()
}

func (daisy *Daisy) Yesterday() *Daisy {
	return New(daisy.Time.AddDate(0, 0, -1)).BeginningOfDay()
}

func (daisy *Daisy) PrevYear() *Daisy {
	return New(daisy.Time.AddDate(-1, 0, 0)).BeginningOfDay()
}

func (daisy *Daisy) NextYear() *Daisy {
	return New(daisy.Time.AddDate(1, 0, 0)).BeginningOfDay()
}

func (daisy *Daisy) PrevMonth() *Daisy {
	return New(daisy.Time.AddDate(0, -1, 0)).BeginningOfDay()
}

func (daisy *Daisy) NextMonth() *Daisy {
	return New(daisy.Time.AddDate(0, 1, 0)).BeginningOfDay()
}

func (daisy *Daisy) BeginningOfDay() *Daisy {
	y, m, d := daisy.Date()
	return New(time.Date(y, m, d, 0, 0, 0, 0, daisy.Time.Location()))
}

func (daisy *Daisy) BeginningOfHour() *Daisy {
	y, m, d := daisy.Date()
	return New(time.Date(y, m, d, daisy.Time.Hour(), 0, 0, 0, daisy.Time.Location()))
}

func BeginningOfDay() *Daisy {
	return New(time.Now()).BeginningOfDay()
}

func (daisy *Daisy) AsTime() time.Time {
	return daisy.Time
}
