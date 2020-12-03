// My implementation of Pricenton 2019 COS418 Assignment 1
package parallelSum

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// Sum numbers from channel `nums` and output sum to `out`.
func sumWorker(nums chan int, out chan int) {
	var sum int
	for num := range nums {
		sum += num
	}
	out <- sum
}

func sum(num int, fileName string) int {
	file, err := os.Open(fileName)
	checkError(err)

	reader := bufio.NewReader(file) // Read numbers
	intSlice, err := readInts(reader)
	checkError(err)

	// Make slice of channels
	numChan := make([]chan int, num)
	resChan := make([]chan int, num)
	for i := 0; i < num; i++ {
		numChan[i] = make(chan int)
		resChan[i] = make(chan int)
	}

	// Launch workers
	for i := 0; i < num; i++ {
		go sumWorker(numChan[i], resChan[i])
	}

	// Supply workers with input
	sz := len(intSlice) / num
	for i := 0; i < num-1; i++ {
		go func(i int) {
			partIntSlice := intSlice[i*sz : (i+1)*sz]
			for j := 0; j < sz; j++ { // Sent to sumWorker
				numChan[i] <- partIntSlice[j]
			}
			close(numChan[i])
		}(i)
	}
	go func() { // Last part
		partIntSlice := intSlice[(num-1)*sz:]
		for j := 0; j < sz; j++ {
			numChan[num-1] <- partIntSlice[j]
		}
		close(numChan[num-1])
	}()

	doneSum := make(chan bool, num)
	// Collect the output from workers
	resSlice := make([]int, num)
	for i := 0; i < num; i++ {
		go func(i int) {
			resSlice[i] = <-resChan[i]
			close(resChan[i])
			doneSum <- true
		}(i)
	}

	for i := 0; i < num; i++ {
		<-doneSum
	}

	// Sum the outputs
	res := 0
	for i := 0; i < len(resSlice); i++ {
		res += resSlice[i]
	}
	return res
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
