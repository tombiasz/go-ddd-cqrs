package main

import (
	"context"
	"fmt"
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

		err := rows.Scan(&m.Id, &m.Code, &m.Email, &m.Description,
			&m.Status, &m.Expdays, &m.ActivatedAt, &m.ExpiredAt, &m.UsedAt,
		)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			os.Exit(1)
		}

		c, errs := coupon.New(
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

		if errs != nil {
			fmt.Fprintf(os.Stderr, "Domain errors: %v\n", errs)
			os.Exit(1)
		}

		fmt.Printf("%s", c)
	}

}
