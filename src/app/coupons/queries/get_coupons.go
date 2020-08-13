package queries

import (
	"context"
	"fmt"
	"time"

	"go-coupons/src/app/coupons/db"
)

// TODO: add pagination
// type GetCouponsQuery struct {
// 	page    uint
// 	perPage uint
// }

type GetCouponsQueryHandler struct {
	DbConnectionFactory db.DbConnectionFactory
}

type GetCouponQueryResult struct {
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

func (h *GetCouponsQueryHandler) Query() []GetCouponQueryResult {
	con := h.DbConnectionFactory.GetConnection()
	defer con.Close()

	rows, err := con.Query(
		context.Background(),
		`SELECT id,
       code,
       email,
       description,
       status,
       expdays,
       activatedAt,
       expiredAt,
       usedAt
			FROM coupons
			ORDER BY activatedAt DESC
		`,
	)

	if err != nil {
		panic(fmt.Sprintf("GetCouponsQuery query failed: %v\n", err))
	}

	var list []GetCouponQueryResult

	for rows.Next() {
		var r GetCouponQueryResult

		err := rows.Scan(
			&r.Id,
			&r.Code,
			&r.Email,
			&r.Description,
			&r.Status,
			&r.Expdays,
			&r.ActivatedAt,
			&r.ExpiredAt,
			&r.UsedAt,
		)

		if err != nil {
			panic(fmt.Sprintf("GetCouponsQuery scan failed: %v\n", err))
		}

		list = append(list, r)
	}

	return list
}
