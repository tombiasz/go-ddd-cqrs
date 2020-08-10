package main

import (
	"context"
	"fmt"
	"go-coupons/src/app/coupons/domain"
	"go-coupons/src/app/coupons/domain/coupon"
	"os"
	"time"
)

func main() {
	conn := NewDbConnection(os.Getenv("DATABASE_URL"))
	defer conn.Close(context.Background())

	rows, err := conn.Query(
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

		var couponIdRaw, codeRaw, emailRaw, descRaw, statusRaw string
		var expdaysRaw uint8
		var activatedAtRaw, expiredAtRaw, usedAtRaw *time.Time

		err := rows.Scan(&couponIdRaw, &codeRaw, &emailRaw, &descRaw,
			&statusRaw, &expdaysRaw, &activatedAtRaw, &expiredAtRaw, &usedAtRaw,
		)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			os.Exit(1)
		}

		couponId, errId := coupon.NewCouponId(couponIdRaw)
		code, errCode := coupon.NewCode(codeRaw)
		email, errEmail := coupon.NewEmail(emailRaw)
		desc, errDesc := coupon.NewDescription(descRaw)

		errs := domain.CombineDomainErrors(errId, errCode, errEmail, errDesc)

		if errs != nil {
			fmt.Fprintf(os.Stderr, "Domain errors: %v\n", errs)
			os.Exit(1)
		}

		var status coupon.Status
		switch statusRaw {
		case coupon.ActiveStatus:
			status = coupon.NewActiveStatus(*activatedAtRaw, expdaysRaw)
		case coupon.UsedStatus:
			status = coupon.NewUsedStatus(*usedAtRaw)
		case coupon.ExpiredStatus:
			status = coupon.NewExpiredStatus(*expiredAtRaw)
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
