package time

import "time"

type RealTimeProvider struct{}

func (t *RealTimeProvider) Now() time.Time {
	return time.Now()
}
