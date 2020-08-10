package dbcoupon

import (
	"time"
)

type CouponModel struct {
	CouponId    string
	Code        string
	Email       string
	Desc        string
	Status      string
	Expdays     uint8
	ActivatedAt *time.Time
	ExpiredAt   *time.Time
	UsedAt      *time.Time
}
