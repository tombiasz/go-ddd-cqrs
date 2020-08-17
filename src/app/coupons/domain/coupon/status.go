package coupon

import (
	"go-coupons/src/app/coupons/domain"
	"time"
)

const ActiveStatus = "Active"
const ExpiredStatus = "Expired"
const UsedStatus = "Used"

type Status interface {
	Status() string
	ActivatedAt() time.Time
	ExpiredAt() *time.Time
	UsedAt() *time.Time
	ExpiresInDays() uint8
	IsActive() bool
	IsUsed() bool
	IsExpired(tp domain.TimeProvider) bool
	Expire(tp domain.TimeProvider) Status
	Use(tp domain.TimeProvider) Status
}

type status struct {
	status        string
	expiresInDays uint8
	activatedAt   time.Time
	expiredAt     *time.Time
	usedAt        *time.Time
}

func (s *status) Status() string {
	return s.status
}

func (s *status) ExpiresInDays() uint8 {
	return s.expiresInDays
}

func (s *status) ActivatedAt() time.Time {
	return s.activatedAt
}

func (s *status) ExpiredAt() *time.Time {
	return s.expiredAt
}

func (s *status) UsedAt() *time.Time {
	return s.usedAt
}

func (s *status) IsExpired(timeProvider domain.TimeProvider) bool {
	if s.expiredAt != nil {
		return true
	}

	if s.IsUsed() {
		return false
	}

	var now = timeProvider.Now()
	var expiresAt = s.activatedAt.AddDate(0, 0, int(s.expiresInDays))

	return now.After(expiresAt)
}

func (s *status) IsActive() bool {
	return s.status == ActiveStatus
}

func (s *status) IsUsed() bool {
	return s.status == UsedStatus
}

func (s *status) Use(tp domain.TimeProvider) Status {
	if !s.IsActive() {
		panic("Invalid status state. Only ActiveStatus can be used")
	}

	now := tp.Now()

	return &status{
		status:        UsedStatus,
		expiresInDays: s.expiresInDays,
		activatedAt:   s.activatedAt,
		expiredAt:     nil,
		usedAt:        &now,
	}
}

func (s *status) Expire(tp domain.TimeProvider) Status {
	if !s.IsActive() {
		panic("Invalid status state. Only ActiveStatus can be expired")
	}

	now := tp.Now()

	return &status{
		status:        ExpiredStatus,
		expiresInDays: s.expiresInDays,
		activatedAt:   s.activatedAt,
		expiredAt:     &now,
		usedAt:        nil,
	}
}

func NewStatus(
	statusStr string,
	expiresInDays uint8,
	activatedAt time.Time,
	expiredAt *time.Time,
	usedAt *time.Time,
) Status {
	// TODO:
	// - check that expiredAt and usedAt are not set at the same time
	return &status{
		statusStr,
		expiresInDays,
		activatedAt,
		expiredAt,
		usedAt,
	}
}

func NewActiveStatus(expiresInDays uint8, activatedAt time.Time) Status {
	return &status{
		status:        ActiveStatus,
		expiresInDays: expiresInDays,
		activatedAt:   activatedAt,
		expiredAt:     nil,
		usedAt:        nil,
	}
}

func NewExpiredStatus(
	expiresInDays uint8,
	activatedAt time.Time,
	expiredAt time.Time,
) Status {
	// TODO:
	// consider additional validation
	// activatedAt must be before expiredAt
	// expiredAt must equal activatedAt + expireInDays
	return &status{
		status:        ExpiredStatus,
		expiresInDays: expiresInDays,
		activatedAt:   activatedAt,
		expiredAt:     &expiredAt,
		usedAt:        nil,
	}
}

func NewUsedStatus(
	expiresInDays uint8,
	activatedAt time.Time,
	usedAt time.Time,
) Status {
	// TODO:
	// consider additional validation
	// activatedAt must be before usedAt
	// usedAt must be less than activatedAt + expireInDays
	return &status{
		status:        UsedStatus,
		expiresInDays: expiresInDays,
		activatedAt:   activatedAt,
		expiredAt:     nil,
		usedAt:        &usedAt,
	}
}
