package coupon

type Coupon struct {
	Id          string
	Email       string
	Code        string
	Description string
	Status      string
}

func New(id, email, code, description, status string) *Coupon {
	return &Coupon{id, email, code, description, status}
}
