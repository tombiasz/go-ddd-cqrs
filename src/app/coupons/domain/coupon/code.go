package coupon

import (
	"fmt"
	"strings"

	"github.com/teris-io/shortid"

	"go-coupons/src/app/coupons/domain"
)

type Code struct {
	value string
}

func NewCode(code string) (*Code, *domain.DomainError) {
	if code == "" {
		return nil, CodeCannotBeEmptyErr
	}

	if len(code) < 8 || len(code) > 11 {
		return nil, CodeIsInvalidErr
	}

	return &Code{code}, nil
}

func GenerateCode() *Code {
	var c = shortid.MustGenerate()

	return &Code{c}
}

func (c *Code) Value() string {
	return strings.ToLower(c.value)
}

func (c *Code) String() string {
	return fmt.Sprintf("<Code: %s >", c.value)
}
