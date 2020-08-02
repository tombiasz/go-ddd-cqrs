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
		var past = time.Now().AddDate(0, 0, -1*expiresInDays).Add(-1 * time.Second)
		var pastTimeProvider = &timeutils.FixedTimeProvider{past}

		a := CreateActiveStatus(expiresInDays, pastTimeProvider)

		e := a.Expire()

		if e.Status() != ExpiredStatus {
			t.Errorf("got %q, want %q", e.Status(), UsedStatus)
		}

		var expectedTime = past.AddDate(0, 0, expiresInDays)
		if !e.expiredAt.Equal(expectedTime) {
			t.Errorf("got %q, want %q", e.expiredAt, expectedTime)
		}
	})
}

func TestActiveStatusIsExpired(t *testing.T) {
	t.Run("return true if active status expired", func(t *testing.T) {
		const expiresInDays = 7
		var now = time.Now()
		var past = now.AddDate(0, 0, -1*expiresInDays).Add(-1 * time.Second)
		var pastTimeProvider = &timeutils.FixedTimeProvider{past}
		var nowTimeProvider = &timeutils.FixedTimeProvider{now}

		a := CreateActiveStatus(expiresInDays, pastTimeProvider)

		isExpired := a.isExpired(nowTimeProvider)

		if !isExpired {
			t.Errorf("expected status to be expire but it did not")
		}
	})

	t.Run("return false if active status not yet expired", func(t *testing.T) {
		const expiresInDays = 7
		var now = time.Now()
		var nowTimeProvider = &timeutils.FixedTimeProvider{now}

		a := CreateActiveStatus(expiresInDays, nowTimeProvider)

		isExpired := a.isExpired(nowTimeProvider)

		if isExpired {
			t.Errorf("expected status to not expire but it did")
		}
	})
}
