package main

import (
	"fmt"

	"go-coupons/src/domain/coupon"
)

func main() {
	c1 := coupon.New("1", "Foo", "Bar", "test", "test")

	fmt.Println(c1)
	fmt.Println(c1.Id())
	fmt.Println(c1.Email())
	fmt.Println(c1.Description())
	fmt.Println(c1.Code())
	fmt.Println(c1.Status())
}
