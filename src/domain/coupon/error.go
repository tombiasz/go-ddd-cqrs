package coupon

import "go-coupons/src/domain"

var DescriptionCannotBeEmptyErr = domain.NewDomainError("description", "description cannot be empty")

var DescriptionCannotBeLongerThan200Chars = domain.NewDomainError("description", "description must have less than 200 characters")
