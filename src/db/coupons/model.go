package dbcoupon

import (
	"time"
)

type CouponModel struct {
	Id          string
	Code        string
	Email       string
	Description string
	Status      string
	Expdays     uint8
	ActivatedAt time.Time
	ExpiredAt   *time.Time
	UsedAt      *time.Time
}
