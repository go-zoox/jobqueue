package jobqueue

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-zoox/uuid"
)

// Worker is a worker for job processing
type Worker struct {
	id        string
	done      *sync.WaitGroup
	readyPool chan chan Job
	jobQueue  chan Job
	quit      chan bool
}

// NewWorker creates a new worker
func NewWorker(readyPool chan chan Job, done *sync.WaitGroup) *Worker {
	worker := &Worker{
		id:        uuid.V4(),
		done:      done,
		readyPool: readyPool,
		jobQueue:  make(chan Job),
		quit:      make(chan bool),
	}

	return worker
}

// Start starts the worker
func (w *Worker) Start() {
	w.done.Add(1)

	go func() {
		for {
			w.readyPool <- w.jobQueue

			select {
			// wait the job in
			case job := <-w.jobQueue:
				func() {
					defer func() {
						if err := recover(); err != nil {
							switch v := err.(type) {
							case string:
								job.Status(JobStatusFailed, errors.New(v))
							case error:
								job.Status(JobStatusFailed, v)
							default:
								job.Status(JobStatusFailed, fmt.Errorf("unknown error: %v", err))
							}
						}
					}()

					job.Status(JobStatusRunning, nil)
					job.Process()
					job.Status(JobStatusDone, nil)
				}()
			case <-w.quit:
				w.done.Done()
				return
			}
		}
	}()
}

// Stop stops the worker
func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
