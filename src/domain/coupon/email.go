package coupon

import "go-coupons/src/domain"

type email struct {
	address string
}

func newEmail(address string) *email {
	return &email{address}
}

func CreateEmail(address string) (*email, *domain.DomainError) {
	if address == "" {
		return nil, domain.NewDomainError("email", "email address cannot be empty")
	}

	if len(address) < 3 || len(address) > 255 {
		return nil, domain.NewDomainError("email", "email address is invalid")
	}

	return &email{address}, nil
}
