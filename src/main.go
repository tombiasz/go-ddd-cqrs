package main

import (
	"fmt"
	"os"

	"go-coupons/src/domain/coupon"
)

func main() {
	c1, err := coupon.Create("1", "Fo", "Bar", "test", "test")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(c1)
	fmt.Println(c1.Id())
	fmt.Println(c1.Email())
	fmt.Println(c1.Description())
	fmt.Println(c1.Code())
	fmt.Println(c1.Status())
}
