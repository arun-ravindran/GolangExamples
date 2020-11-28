package counter


type Counter struct {
	count int
	updateRequests chan *UpdateOp
}

type UpdateOp struct {
	amount int
	confirm chan int
}


func NewCounter() *Counter {
	return &Counter {
		count: 0,
		updateRequests: make(chan *UpdateOp, 4),
	}
}

// Update requests, creates a UpdateOp object that is inserted into the channel
// When the update is processed, a confirm value is sent back on the channel contained in the request
func (ctr *Counter) Update(amt int) int {
	uOp := &UpdateOp{amount: amt, confirm: make(chan int)}
	ctr.updateRequests <- uOp
	newVal :=  <-uOp.confirm
	return newVal
}

// Process each request on the request channel, and sent cofirmation back on the channel of each request
func (ctr *Counter) ProcessUpdates() {
	for {
		select {
			case request := <-ctr.updateRequests:
				ctr.count += request.amount
				request.confirm <- ctr.count
		}
	}
}
