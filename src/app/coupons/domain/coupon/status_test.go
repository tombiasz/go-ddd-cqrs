package coupon

import (
	timeutils "go-coupons/src/utils/time"
	"testing"
	"time"
)

func TestCreateActiveStatus(t *testing.T) {
	t.Run("creates an active status", func(t *testing.T) {
		var fakeNow = time.Now()
		var fixedTimeProvider = &timeutils.FixedTimeProvider{fakeNow}
		const expiresInDays = 7

		s := CreateActiveStatus(expiresInDays, fixedTimeProvider)

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
		var fakeNow = time.Now()
		var fixedTimeProvider = &timeutils.FixedTimeProvider{fakeNow}

		a := CreateActiveStatus(expiresInDays, fixedTimeProvider)

		u := a.Use(fixedTimeProvider)

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
		var fakeNow = time.Now()
		var fixedTimeProvider = &timeutils.FixedTimeProvider{fakeNow}

		a := CreateActiveStatus(expiresInDays, fixedTimeProvider)

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
