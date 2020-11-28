// Implements pool of worker threads synchronized with Go channels

/*
	tasks are buffered in taskChannel of capacity corresponding to number of worker threads
	AddTask - inserts task in taskChannel
	Run - runs tasks from taskChannel, and inserts errors into errChannel
	Results - returns errChannel
	
*/
package workerpool

import (
	"errors"
)

var ErrIllegalSize = errors.New("Illegal worker pool size")
var ErrMaxTasks = errors.New("Exceeded maximum number of tasks")
var ErrTask = errors.New("Task error")

type WorkerPool struct {
	taskChannel	chan Task
	errChannel chan error
}

type Task func() error


func NewWorkerPool(size int) (*WorkerPool, error) {
	if size <= 0 {
		return nil, ErrIllegalSize
	}
	return &WorkerPool{taskChannel: make(chan Task, size), errChannel: make(chan error, size)}, nil
}

func (wp *WorkerPool) Run() {
	for task := range wp.taskChannel {
		go func() {
			err := task() // Execute task
			wp.errChannel<- err
		}()
	}
}

func (wp *WorkerPool) AddTask(task Task) error {
	select {
		case wp.taskChannel<- task:
		default: return ErrMaxTasks
	}
	return nil
}

func (wp *WorkerPool) Results() <-chan error  {
	close (wp.taskChannel)
	return wp.errChannel
}

