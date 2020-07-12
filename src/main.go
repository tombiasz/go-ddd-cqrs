package main

import (
	"fmt"
	"os"

	"go-coupons/src/domain/coupon"
)

func main() {
	// sad path
	_, err := coupon.Create("1", "Fo", "Bar", "", "test")

	if err != nil {
		// print all
		fmt.Println(err)

		for _, r := range err {
			// print one by one one
			fmt.Println(r)
		}

		for _, r := range err {
			// print every property
			fmt.Println(r.Property)
			fmt.Println(r.Message)
		}
	}

	// happy path
	c1, err1 := coupon.Create("1", "Foo", "Bar", "Description", "test")

	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}

	fmt.Println(c1)
	fmt.Println(c1.Id())
	fmt.Println(c1.Email())
	fmt.Println(c1.Description())
	fmt.Println(c1.Code())
	fmt.Println(c1.Status())
}
