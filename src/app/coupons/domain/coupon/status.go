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
}

type activeStatus struct {
	activatedAt   time.Time
	expiresInDays uint8
}

func NewActiveStatus(activatedAt time.Time, expiresInDays uint8) *activeStatus {
	return &activeStatus{activatedAt, expiresInDays}
}

func CreateActiveStatus(expiresInDays uint8, timeProvider domain.TimeProvider) *activeStatus {
	return &activeStatus{
		activatedAt:   timeProvider.Now(),
		expiresInDays: expiresInDays,
	}
}

func (s *activeStatus) Status() string {
	return ActiveStatus
}

func (s *activeStatus) isExpired(timeProvider domain.TimeProvider) bool {
	var expiresAt = s.activatedAt.AddDate(0, 0, int(s.expiresInDays))
	var now = timeProvider.Now()

	return now.After(expiresAt)
}

func (s *activeStatus) Expire() *expiredStatus {
	var expiredAt = s.activatedAt.AddDate(0, 0, int(s.expiresInDays))

	return &expiredStatus{expiredAt}
}

func (s *activeStatus) Use(timeProvider domain.TimeProvider) *usedStatus {
	return &usedStatus{
		usedAt: timeProvider.Now(),
	}
}

type expiredStatus struct {
	expiredAt time.Time
}

func NewExpiredStatus(expiredAt time.Time) *expiredStatus {
	return &expiredStatus{expiredAt}
}

func (s *expiredStatus) Status() string {
	return ExpiredStatus
}

type usedStatus struct {
	usedAt time.Time
}

func NewUsedStatus(usedAt time.Time) *usedStatus {
	return &usedStatus{usedAt}
}

func (s *usedStatus) Status() string {
	return UsedStatus
}
