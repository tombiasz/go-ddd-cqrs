package identity

import (
	"go-coupons/src/domain"
)

type FixedIdentityProvider struct {
	Value string
}

func (p *FixedIdentityProvider) NextID() *domain.ID {
	return domain.CreateID(p.Value)
}
