package db

import (
	"context"
	"fmt"
	"go-coupons/src/app/coupons/domain"
	"go-coupons/src/app/coupons/domain/coupon"
	"time"
)

// TODO: move to separate file
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
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO UPDATE SET
			email = EXCLUDED.email,
			code = EXCLUDED.code,
			description = EXCLUDED.description,
			status = EXCLUDED.status,
			expdays = EXCLUDED.expdays,
			activatedat = EXCLUDED.activatedat,
			expiredat = EXCLUDED.expiredat,
			usedat = EXCLUDED.usedat
		`,
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

func (r *couponRepository) GetCouponByEmailAndCode(email *coupon.Email, code *coupon.Code) (*coupon.Coupon, *domain.DomainError) {
	con := r.dbConnectionFactory.GetConnection()

	row := con.QueryRow(
		context.Background(),
		`SELECT
			id,
			email,
			code,
			description,
			status,
			expdays,
			activatedat,
			expiredat,
			usedat
		 FROM coupons
			WHERE email = $1
				AND code = $2`,
		email.Address(),
		code.Value(),
	)

	var m CouponModel

	err := row.Scan(
		&m.Id,
		&m.Email,
		&m.Code,
		&m.Description,
		&m.Status,
		&m.Expdays,
		&m.ActivatedAt,
		&m.ExpiredAt,
		&m.UsedAt,
	)

	if err != nil {
		panic(fmt.Sprintf("GetCouponByEmailAndCode scan failed: %v\n", err))
	}

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
		panic(fmt.Sprintf("GetCouponByEmailAndCode converting model to entity failed: %v\n", errDomain))
	}

	return c, nil
}

func (r *couponRepository) GetExpiredCoupons() ([]*coupon.Coupon, *domain.DomainError) {
	con := r.dbConnectionFactory.GetConnection()

	rows, err := con.Query(
		context.Background(),
		`SELECT
			id,
			email,
			code,
			description,
			status,
			expdays,
			activatedat,
			expiredat,
			usedat
		 FROM coupons
			WHERE activatedat + interval '1 day' * expdays < now()
				AND status = 'Active'
		`,
	)

	if err != nil {
		panic(fmt.Sprintf("GetExpiredCoupons query failed: %v\n", err))
	}

	var coupons []*coupon.Coupon

	for rows.Next() {
		var m CouponModel

		err := rows.Scan(
			&m.Id,
			&m.Email,
			&m.Code,
			&m.Description,
			&m.Status,
			&m.Expdays,
			&m.ActivatedAt,
			&m.ExpiredAt,
			&m.UsedAt,
		)

		if err != nil {
			panic(fmt.Sprintf("GetExpiredCoupons scan failed: %v\n", err))
		}

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
			panic(fmt.Sprintf("GetExpiredCoupons converting model to entity failed: %v\n", errDomain))
		}

		coupons = append(coupons, c)
	}

	// TODO: are two return values needed?
	return coupons, nil
}
