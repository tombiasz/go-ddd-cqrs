package coupon

import (
	"fmt"
	"go-coupons/src/app/coupons/domain"

	uuid "github.com/google/uuid"
)

type couponId struct {
	value uuid.UUID
}

// TODO: missing test
func NewCouponId(value string) (*couponId, *domain.DomainError) {
	id, err := uuid.Parse(value)

	if err != nil {
		return nil, CouponIdIsInvalidErr
	}

	return &couponId{id}, nil
}

func GenerateCouponId() *couponId {
	var id = uuid.New()

	return &couponId{id}
}

func (id *couponId) Value() string {
	return id.value.String()
}

func (id *couponId) String() string {
	return fmt.Sprintf("<CouponID: %s >", id.value)
}
