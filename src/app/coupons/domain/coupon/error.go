package coupon

import "go-coupons/src/app/coupons/domain"

var DescriptionCannotBeEmptyErr = domain.NewDomainError(
	"coupons.description.empty",
	"description cannot be empty",
)

var DescriptionCannotBeLongerThan200CharsErr = domain.NewDomainError(
	"coupons.description.max-length",
	"description must have less than 200 characters",
)

var EmailCannotBeEmptyErr = domain.NewDomainError(
	"coupons.email.empty",
	"email address cannot be empty",
)

var EmailIsInvalidErr = domain.NewDomainError(
	"coupons.email.invalid",
	"email address is invalid",
)

var CouponAlreadyUsedErr = domain.NewDomainError(
	"coupons.coupon.is-used",
	"coupon already been used",
)

var CodeCannotBeEmptyErr = domain.NewDomainError(
	"coupons.code.empty",
	"code cannot be empty",
)

var CodeIsInvalidErr = domain.NewDomainError(
	"coupons.code.invalid",
	"code is invalid",
)

var CouponCannotBeNotExpiredErr = domain.NewDomainError(
	"coupons.status.cannot-be-marked-as-expired",
	"coupon cannot be marked as expired",
)

var CouponIdIsInvalidErr = domain.NewDomainError(
	"coupons.couponId.invalid",
	"coupon id is invalid",
)
