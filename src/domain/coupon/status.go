package coupon

import (
	"go-coupons/src/domain"
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

func CreateActiveStatus(expiresInDays uint8, timeProvider domain.TimeProvider) *activeStatus {
	return &activeStatus{
		activatedAt:   timeProvider.Now(),
		expiresInDays: expiresInDays,
	}
}

func (s *activeStatus) Status() string {
	return ActiveStatus
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

func (s *expiredStatus) Status() string {
	return ExpiredStatus
}

type usedStatus struct {
	usedAt time.Time
}

func (s *usedStatus) Status() string {
	return UsedStatus
}
