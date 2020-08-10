package main

import (
	"context"
	"fmt"
	"go-coupons/src/app/coupons/domain"
	"go-coupons/src/app/coupons/domain/coupon"
	dbcoupon "go-coupons/src/db/coupons"
	"os"
)

func main() {
	con := NewDbConnection(os.Getenv("DATABASE_URL"))
	defer con.Close()

	rows, err := con.Query(
		context.Background(),
		`select
			id,  code, email, description, status, expdays, activatedAt, expiredAt, usedAt
		from coupons`,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		var m dbcoupon.CouponModel

		err := rows.Scan(&m.CouponId, &m.Code, &m.Email, &m.Desc,
			&m.Status, &m.Expdays, &m.ActivatedAt, &m.ExpiredAt, &m.UsedAt,
		)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			os.Exit(1)
		}

		couponId, errId := coupon.NewCouponId(m.CouponId)
		code, errCode := coupon.NewCode(m.Code)
		email, errEmail := coupon.NewEmail(m.Email)
		desc, errDesc := coupon.NewDescription(m.Desc)

		errs := domain.CombineDomainErrors(errId, errCode, errEmail, errDesc)

		if errs != nil {
			fmt.Fprintf(os.Stderr, "Domain errors: %v\n", errs)
			os.Exit(1)
		}

		var status coupon.Status
		switch m.Status {
		case coupon.ActiveStatus:
			status = coupon.NewActiveStatus(*m.ActivatedAt, m.Expdays)
		case coupon.UsedStatus:
			status = coupon.NewUsedStatus(*m.UsedAt)
		case coupon.ExpiredStatus:
			status = coupon.NewExpiredStatus(*m.ExpiredAt)
		}

		coupon := coupon.New(
			couponId,
			email,
			code,
			desc,
			status,
		)

		fmt.Printf("%s", coupon)
	}

}
