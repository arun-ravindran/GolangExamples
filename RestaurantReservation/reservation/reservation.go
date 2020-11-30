package reservation

import (
	"RestaurantReservation/customer"
	"RestaurantReservation/restaurant"
)

type Reservation interface {
}

type reservation struct {
	id         int
	customer   customer.Customer
	restaurant restaurant.Restaurant
}

func NewReservation(id int, c customer.Customer, r restaurant.Restaurant, sizeOfParty int) Reservation {
	r.ReserveTable(sizeOfParty)
	return &reservation{
		id:         id,
		customer:   c,
		restaurant: r,
	}
}
