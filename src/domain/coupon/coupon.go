package coupon

type coupon struct {
	Id          string
	Email       string
	Code        string
	Description string
	Status      string
}

func New(id, email, code, description, status string) *coupon {
	return &coupon{id, email, code, description, status}
}
