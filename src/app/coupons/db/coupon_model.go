package db

import (
	"fmt"
	"go-coupons/src/app/coupons/domain/coupon"
	"time"
)

type couponModel struct {
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

func (m *couponModel) ToEntity() *coupon.Coupon {
	c, errDomain := coupon.New(
		m.Id,
		m.Email,
		m.Code,
		m.Description,
		m.Status,
		m.Expdays,
		m.ActivatedAt,
		m.ExpiredAt,
		m.UsedAt,
	)

	if errDomain != nil {
		panic(fmt.Sprintf("Converting model to entity failed: %v\n", errDomain))
	}

	return c
}
