package time

import "time"

type FixedTimeProvider struct {
	Time time.Time
}

func (t *FixedTimeProvider) Now() time.Time {
	return t.Time
}
