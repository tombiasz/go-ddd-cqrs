package main

import (
	"fmt"
	"go-coupons/src/domain"

	"go-coupons/src/domain/coupon"
)

func main() {
	_, err1 := coupon.CreateEmail("Fo")
	_, err2 := coupon.CreateDescription("")

	err := domain.CombineDomainErrors(err1, err2)

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

	_, err3 := coupon.Create("1", nil, "Bar", nil, "test")

	if err3 != nil {
		fmt.Println(err3)
	}

	e2, _ := coupon.CreateEmail("Foo")
	d2, _ := coupon.CreateDescription("description")

	// happy path
	c1, _ := coupon.Create("1", e2, "Bar", d2, "test")

	fmt.Println(c1)
	fmt.Println(c1.Id())
	fmt.Println(c1.Email())
	fmt.Println(c1.Description())
	fmt.Println(c1.Code())
	fmt.Println(c1.Status())
}
