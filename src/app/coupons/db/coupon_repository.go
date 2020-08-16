package db

import (
	"context"
	"fmt"
	"go-coupons/src/app/coupons/domain"
	"go-coupons/src/app/coupons/domain/coupon"
)

type couponRepository struct {
	dbConnectionFactory DbConnectionFactory
}

func NewCouponRepository(dbConnectionFactory DbConnectionFactory) *couponRepository {
	return &couponRepository{dbConnectionFactory}
}

func (r *couponRepository) Save(c *coupon.Coupon) *domain.DomainError {
	con := r.dbConnectionFactory.GetConnection()

	_, err := con.Exec(
		context.Background(),
		`INSERT INTO coupons (
			id,
			email,
			code,
			description,
			status,
			expdays,
			activatedat,
			expiredat,
			usedat
		) VALUES ($1, $2,	$3, $4, $5, $6, $7, $8, $9)`,
		c.Id(),
		c.Email(),
		c.Code(),
		c.Description(),
		c.Status(),
		c.ExpireInDays(),
		c.ActivatedAt(),
		c.ExpiredAt(),
		c.UsedAt(),
	)

	if err != nil {
		panic(fmt.Sprintf("Save coupon failed: %v\n", err))
	}

	return nil
}

func (r *couponRepository) GetCouponByEmailAndCode(*coupon.Email, *coupon.Code) (*coupon.Coupon, *domain.DomainError) {
	return nil, nil
}

func (r *couponRepository) GetExpiredCoupons() ([]*coupon.Coupon, *domain.DomainError) {
	return nil, nil
}
