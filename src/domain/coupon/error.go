package coupon

import "go-coupons/src/domain"

var DescriptionCannotBeEmptyErr = domain.NewDomainError("description", "description cannot be empty")

var DescriptionCannotBeLongerThan200CharsErr = domain.NewDomainError("description", "description must have less than 200 characters")

var EmailCannotBeEmptyErr = domain.NewDomainError("email", "email address cannot be empty")

var EmailIsInvalidErr = domain.NewDomainError("email", "email address is invalid")
