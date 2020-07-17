package domain

type ID struct {
	value string
}

func CreateID(value string) *ID {
	return &ID{value}
}

type IdentityProvider interface {
	NextID() *ID
}

func (id *ID) Value() string {
	return id.value
}
