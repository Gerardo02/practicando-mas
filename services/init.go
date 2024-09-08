package services

type Services struct {
	Order *Order
	OAuth *OAuth
}

func NewServices() *Services {
	return &Services{
		Order: NewOrder(),
		OAuth: NewAuth(),
	}
}
