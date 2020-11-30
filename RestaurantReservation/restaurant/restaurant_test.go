package restaurant

import "testing"

func TestReserveTableUnderCapacity(t *testing.T) {
	myRest := NewRestaurant(1, "Midwood", "BBQ", 10)
	if myRest.ReserveTable(5) != nil {
		t.Error("TestResrveTable failed")
	}
}

func TestReleaseTableUnderOccupied(t *testing.T) {
	myRest := NewRestaurant(1, "Midwood", "BBQ", 10)
	myRest.ReserveTable(5)
	if myRest.ReleaseTable(5) != nil {
		t.Error("TestReleaseTable failed")
	}
}
