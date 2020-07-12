package coupon

import "go-coupons/src/domain"

type email struct {
	value string
}

func newEmail(value string) *email {
	return &email{value}
}

func createEmail(value string) (*email, *domain.DomainError) {
	if value == "" {
		return nil, domain.NewDomainError("email", "email cannot be empty")
	}

	if len(value) < 3 || len(value) > 255 {
		return nil, domain.NewDomainError("email", "email is invalid")
	}

	return &email{value}, nil
}
