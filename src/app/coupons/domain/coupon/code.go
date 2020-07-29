package coupon

import (
	"fmt"

	"github.com/teris-io/shortid"

	"go-coupons/src/app/coupons/domain"
)

type Code struct {
	value string
}

// TODO: add validation + tests
func NewCode(c string) (*Code, *domain.DomainError) {
	return &Code{c}, nil
}

func GenerateCode() *Code {
	var c = shortid.MustGenerate()

	return &Code{c}
}

func (c *Code) Value() string {
	return c.value
}

func (c *Code) String() string {
	return fmt.Sprintf("<Code: %s >", c.value)
}
