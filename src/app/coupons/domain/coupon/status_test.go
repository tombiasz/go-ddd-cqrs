package coupon

import (
	timeutils "go-coupons/src/utils/time"
	"testing"
	"time"
)

func TestIsExpired(t *testing.T) {
	const expiresInDays = 7
	var now = time.Now()
	var fixedTimeProvider = &timeutils.FixedTimeProvider{now}
	var twoDaysAgo = time.Now().AddDate(0, 0, -1*2).Add(-1 * time.Second)
	var sevenDaysAgo = time.Now().AddDate(0, 0, -1*expiresInDays).Add(-1 * time.Second)

	t.Run("active status should not be expired when not expiresInDays have passed", func(t *testing.T) {
		s := NewActiveStatus(expiresInDays, twoDaysAgo)

		isExpired := s.IsExpired(fixedTimeProvider)

		if isExpired {
			t.Errorf("status should no be expired")
		}
	})

	t.Run("active status should be expired when expiresInDays have passed", func(t *testing.T) {
		s := NewActiveStatus(expiresInDays, sevenDaysAgo)

		isExpired := s.IsExpired(fixedTimeProvider)

		if !isExpired {
			t.Errorf("status should be expired")
		}
	})

	t.Run("used status should not be expired", func(t *testing.T) {
		s := NewUsedStatus(expiresInDays, sevenDaysAgo, now)

		isExpired := s.IsExpired(fixedTimeProvider)

		if isExpired {
			t.Errorf("status should be expired")
		}
	})

	t.Run("expired status should be expired", func(t *testing.T) {
		s := NewExpiredStatus(expiresInDays, sevenDaysAgo, now)

		isExpired := s.IsExpired(fixedTimeProvider)

		if !isExpired {
			t.Errorf("status should not be expired")
		}
	})
}

func TestUse(t *testing.T) {
	const expiresInDays = 7
	var now = time.Now()
	var fixedTimeProvider = &timeutils.FixedTimeProvider{now}
	var twoDaysAgo = time.Now().AddDate(0, 0, -1*2).Add(-1 * time.Second)
	var sevenDaysAgo = time.Now().AddDate(0, 0, -1*expiresInDays).Add(-1 * time.Second)

	t.Run("active status can be used", func(t *testing.T) {
		active := NewActiveStatus(expiresInDays, twoDaysAgo)

		used := active.Use(fixedTimeProvider)

		if !used.IsUsed() {
			t.Errorf("invalid status; expected UseStatus but got %s", used.Status())
		}

		if !used.UsedAt().Equal(now) {
			t.Errorf("invalid usedAt; expected %s but got %s", now, used.UsedAt())
		}
	})

	t.Run("used status cannot be used ", func(t *testing.T) {
		used := NewUsedStatus(expiresInDays, twoDaysAgo, now)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		_ = used.Use(fixedTimeProvider)
	})

	t.Run("expired status cannot be used ", func(t *testing.T) {
		expired := NewExpiredStatus(expiresInDays, sevenDaysAgo, now)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		_ = expired.Use(fixedTimeProvider)
	})
}

func TestExpire(t *testing.T) {
	const expiresInDays = 7
	var now = time.Now()
	var fixedTimeProvider = &timeutils.FixedTimeProvider{now}
	var twoDaysAgo = time.Now().AddDate(0, 0, -1*2).Add(-1 * time.Second)
	var sevenDaysAgo = time.Now().AddDate(0, 0, -1*expiresInDays).Add(-1 * time.Second)

	t.Run("active status can be expired ", func(t *testing.T) {
		active := NewActiveStatus(expiresInDays, twoDaysAgo)

		expired := active.Expire(fixedTimeProvider)

		if !expired.IsExpired(fixedTimeProvider) {
			t.Errorf("invalid status; expected ExpiredStatus but got %s", expired.Status())
		}

		if !expired.ExpiredAt().Equal(now) {
			t.Errorf("invalid expiredAt; expected %s but got %s", now, expired.UsedAt())
		}
	})

	t.Run("used status cannot be expired", func(t *testing.T) {
		used := NewUsedStatus(expiresInDays, twoDaysAgo, now)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		_ = used.Expire(fixedTimeProvider)
	})

	t.Run("expired status cannot be expired", func(t *testing.T) {
		expired := NewExpiredStatus(expiresInDays, sevenDaysAgo, now)

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		_ = expired.Use(fixedTimeProvider)
	})
}
