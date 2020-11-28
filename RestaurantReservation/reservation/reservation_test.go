package reservation

import (
	"RestaurantReservation/restaurant"
	"RestaurantReservation/customer"
)


import  "testing"

func TestNewReservation(t *testing.T) {
	rest := restaurant.NewRestaurant(1, "Smokehouse", "BBQ", 10)
	cust := customer.NewCustomer(1, 7706)
	_ = NewReservation(1, cust, rest, 5)
}
