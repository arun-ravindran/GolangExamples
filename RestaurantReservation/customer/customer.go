package customer

type Customer interface {
}

type customer struct {
	customerID int
	phoneNumber int
}

func NewCustomer(id int, phone int) Customer {
	return &customer{customerID:id, phoneNumber:phone}
}



