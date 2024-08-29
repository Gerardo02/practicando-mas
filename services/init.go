package services

type Services struct {
	Order *Order
}

func NewServices() *Services {
	return &Services{
		Order: NewOrder(),
	}
}
