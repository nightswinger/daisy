package daisy

import "time"

type Daisy struct {
	time.Time
}

func New(t time.Time) *Daisy {
	return &Daisy{t}
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
