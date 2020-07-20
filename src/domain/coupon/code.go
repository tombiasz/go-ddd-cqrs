package coupon

import (
	"fmt"

	"github.com/teris-io/shortid"
)

type code struct {
	value string
}

func GenerateCode() *code {
	var c = shortid.MustGenerate()

	return &code{c}
}

func (c *code) Value() string {
	return c.value
}

func (c *code) String() string {
	return fmt.Sprintf("<Code: %s >", c.value)
}
