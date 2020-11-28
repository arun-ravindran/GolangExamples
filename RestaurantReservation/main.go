package main

import (
	"RestaurantReservation/reservation"
	"RestaurantReservation/customer"
	"RestaurantReservation/restaurant"
)

func main() {
	customerDB := make(map[int]*customer.Customer)
	restaurantDB := make(map[int]*restaurant.Restaurant)

	// Insert customer entries - from external database
	customer1 := customer.NewCustomer(1, 7706)
	customer2 := customer.NewCustomer(2, 8427)
	customerDB[1] = &customer1
	customerDB[2] = &customer2


	// Insert restaurant entries - from external database
	restaurant1 := restaurant.NewRestaurant(1, "Smokehouse", "BBQ", 10)
	restaurant2 := restaurant.NewRestaurant(2, "Azteca", "Mexcian", 20)
	restaurantDB[1] = &restaurant1
	restaurantDB[2] = &restaurant2


	// Make new reservation
	_ = reservation.NewReservation(1, customer1, restaurant1, 4)

}



