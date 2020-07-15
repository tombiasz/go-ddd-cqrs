package coupon

import (
	"testing"
	"time"
)

var fakeNow = time.Now()

type fakeTimeProvider struct{}

func (t *fakeTimeProvider) Now() time.Time {
	return fakeNow
}

func TestCreateActiveStatus(t *testing.T) {
	t.Run("creates an active status", func(t *testing.T) {
		const expiresInDays = 7

		s := CreateActiveStatus(expiresInDays, &fakeTimeProvider{})

		if s.Status() != ActiveStatus {
			t.Errorf("got %q, want %q", s.Status(), ActiveStatus)
		}

		if !s.activatedAt.Equal(fakeNow) {
			t.Errorf("got %q, want %q", s.activatedAt, fakeNow)
		}

		if s.expiresInDays != expiresInDays {
			t.Errorf("got %q, want %q", s.expiresInDays, expiresInDays)
		}
	})
}

func TestActiveStatusUse(t *testing.T) {
	t.Run("using active status should return used status", func(t *testing.T) {
		const expiresInDays = 7
		a := CreateActiveStatus(expiresInDays, &fakeTimeProvider{})

		u := a.Use(&fakeTimeProvider{})

		if u.Status() != UsedStatus {
			t.Errorf("got %q, want %q", u.Status(), UsedStatus)
		}

		if !u.usedAt.Equal(fakeNow) {
			t.Errorf("got %q, want %q", u.usedAt, fakeNow)
		}
	})
}

func TestActiveStatusExpire(t *testing.T) {
	t.Run("expiring active status should return expired status", func(t *testing.T) {
		const expiresInDays = 7
		a := CreateActiveStatus(expiresInDays, &fakeTimeProvider{})

		e := a.Expire()

		if e.Status() != ExpiredStatus {
			t.Errorf("got %q, want %q", e.Status(), UsedStatus)
		}

		var expectedTime = fakeNow.AddDate(0, 0, expiresInDays)
		if !e.expiredAt.Equal(expectedTime) {
			t.Errorf("got %q, want %q", e.expiredAt, expectedTime)
		}
	})
}
