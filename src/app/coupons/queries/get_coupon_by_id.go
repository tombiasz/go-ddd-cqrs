package queries

import (
	"context"
	"fmt"
	"time"

	"go-coupons/src/app/coupons/db"

	"github.com/jackc/pgx/v4"
)

type GetCouponByIdQueryHandler struct {
	DbConnectionFactory db.DbConnectionFactory
}

type GetCouponByIdQueryResult struct {
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

func (h *GetCouponByIdQueryHandler) Query(couponId string) *GetCouponByIdQueryResult {
	con := h.DbConnectionFactory.GetConnection()
	defer con.Close()

	row := con.QueryRow(
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
			WHERE id = $1
		`,
		couponId,
	)

	var r GetCouponByIdQueryResult

	err := row.Scan(
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
		if err == pgx.ErrNoRows {
			return nil
		}

		panic(fmt.Sprintf("GetCouponByIdQuery scan failed: %v\n", err))
	}

	return &r
}
