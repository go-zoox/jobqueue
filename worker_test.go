package jobqueue

import (
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	taskDone := false

	pool := make(chan chan Job)
	wg := &sync.WaitGroup{}

	worker := NewWorker(pool, wg)
	worker.Start()

	workerChannel := <-pool
	workerChannel <- NewJob(func() {
		taskDone = true
	}, func(status int, err error) {
		t.Logf("Job status: %d, err: %v", status, err)
	})

	<-pool

	if !taskDone {
		t.Error("task should be done")
	}
}
