package counter

import "testing"

func  TestNewCounter(t *testing.T) {
	_ = NewCounter()
}

func TestUpdate(t *testing.T) {
	ctr := NewCounter()
	go ctr.ProcessUpdates()
	for  i := 1; i <= 4; i++ {
		go func() {
			newval := ctr.Update(1)
			if newval != i {
				t.Errorf("Incorrect update")
			}
		}()
	}
}



