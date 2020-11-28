# Restaurant reservation - using interfaces in Golang for a decoupled design

Goal: To design an OpenTable like restaurant reservation system.

Key entities

Restaurant
- id
- name
- description
- maxCapacity
- numFree
+ NewRestaurant
+ ReserveTable(num)
+ ReleaseTable(num)

Customer
- id
- phoneNumber

Reservation
- id
- restaurant
- customer
+ NewReservation
+ CancelReservation


main
- Create 2 new restaurants // Read from restaurant table
- Create 2 customers // Write to cutomer table
- Customer 1 reserves first restaurant // Write to reservation table
- Customer 2 reserves second restaurant // Write to reservatation table
