package workerpool

import (
	"testing"
	"time"
)

// Test Constructor
func TestNewWorkerPool(t *testing.T) {
	_, err := NewWorkerPool(4)
		if err != nil {
			t.Error("TestNewWorkerPool failed", err)
		}

}

// Test AddTask
func TestAddTask(t *testing.T) {
	wPool, _ := NewWorkerPool(2)
	task := Task(func() error {return ErrTask})
	err := wPool.AddTask(task)
	if err != nil {
		t.Error("Add task failed")
	}
}


// Test add more tasks than possible to handle
func TestMaxTasks(t *testing.T) {
	wPool, _ := NewWorkerPool(2)
	var err error
	task1 := Task(func() error {return ErrTask})
	if err = wPool.AddTask(task1); err != nil {
		t.Error("Add task failed")
	}
	task2 := Task(func() error {return ErrTask})
	if err = wPool.AddTask(task2); err != nil {
		t.Error("Add task failed")
	}
	task3 := Task(func() error {return ErrTask})
	if err = wPool.AddTask(task3); err !=  ErrMaxTasks {
		t.Error("Add task failed")
	}

}

// Run tasks and check error channel
func TestRunTasks(t *testing.T) {
	wPool, _ := NewWorkerPool(2)
	var err error
	task1 := Task(func() error {return ErrTask})
	if err = wPool.AddTask(task1); err != nil {
		t.Error("Add task failed")
	}
	task2 := Task(func() error {return ErrTask})
	if err = wPool.AddTask(task2); err != nil {
		t.Error("Add task failed")
	}
	var errRes []error
	go func() {
		time.Sleep(100*time.Millisecond) // Wait until workers get to run. 
		wPoolResults := wPool.Results() // Will close channel to allow test to exit
		for i := 0; i < 2; i++ {
			select {
				case err = <-wPoolResults:
					errRes = append(errRes, err)
				default:
					errRes = append(errRes, nil)
			}
		}
	}()
	wPool.Run() // Run workers
	for i := 0; i < len(errRes); i++ {
		if errRes[i] != ErrTask {
			t.Error("Error result failed")
		}
	}
}

