package restaurant

import "errors"

type Restaurant interface {
	ReserveTable(size interface{}) error
	ReleaseTable(size interface{})error
}

type restaurant struct {
	id int
	name string
	description string
	maxCapacity int
	numFree int
}


func NewRestaurant(id int, n string, d string, c int) Restaurant {
	return &restaurant{id: id, name: n, description: d, maxCapacity: c, numFree: c}
}

func (r *restaurant) ReserveTable(size interface{}) error {
	if r.numFree < size.(int) {
		return errors.New("Not enough free tables")
	}
	r.numFree -= size.(int)
	return nil
}

func (r *restaurant) ReleaseTable(size interface{}) error {
	if size.(int) > r.maxCapacity - r.numFree {
		return errors.New("Number exceeds occupied table")
	}
	r.numFree += size.(int)
	return nil
}
