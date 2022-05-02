package jobqueue

import (
	"sync"
)

// JobQueue is a queue for jobs
type JobQueue struct {
	queue     chan Job
	readyPool chan chan Job
	workers   []*Worker
	quit      chan bool
	//
	wgMaster *sync.WaitGroup
	wgWorker *sync.WaitGroup
}

// New creates a new job queue
func New(maxWorkers int) *JobQueue {
	wgWorker := &sync.WaitGroup{}
	wgMaster := &sync.WaitGroup{}
	readyPool := make(chan chan Job, maxWorkers)
	workers := make([]*Worker, maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		workers[i] = NewWorker(readyPool, wgWorker)
	}

	return &JobQueue{
		queue:     make(chan Job),
		readyPool: readyPool,
		workers:   workers,
		quit:      make(chan bool),
		wgWorker:  wgWorker,
		wgMaster:  wgMaster,
	}
}

// AddJob adds a job to the queue
func (q *JobQueue) AddJob(job Job) {
	job.Status(JobStatusPending, nil)

	q.queue <- job
}

// Start starts the job queue
func (q *JobQueue) Start() {
	// 1. workers start
	for i := 0; i < len(q.workers); i++ {
		q.workers[i].Start()
	}

	// 2. main start
	q.wgMaster.Add(1)
	go func() {
		for {
			select {
			case job := <-q.queue:
				workerChannel := <-q.readyPool
				workerChannel <- job
			case <-q.quit:
				for i := 0; i < len(q.workers); i++ {
					q.workers[i].Stop()
				}

				q.wgWorker.Wait()
				q.wgMaster.Done()
				return
			}
		}
	}()
}

// Stop stops the job queue
func (q *JobQueue) Stop() {
	// Stopping queue
	q.quit <- true

	// Stopped queue
	q.wgMaster.Wait()
}
