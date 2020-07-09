package main

import (
	"fmt"

	"go-coupons/src/domain/coupon"
)

func main() {
	c1 := coupon.New("1", "Foo", "Bar", "test", "test")

	fmt.Println(c1)
}
