package coupon

import (
	"fmt"

	uuid "github.com/google/uuid"
)

type couponId struct {
	value uuid.UUID
}

func CreateCouponId() *couponId {
	var id = uuid.New()

	return &couponId{id}
}

func (id *couponId) Value() string {
	return id.value.String()
}

func (id *couponId) String() string {
	return fmt.Sprintf("<CouponID: %s >", id.value)
}
