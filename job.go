package jobqueue

const (
	// JobStatusPending means the job is waiting to be processed.
	JobStatusPending = iota
	// JobStatusRunning means the job is being processed.
	JobStatusRunning
	// JobStatusDone means the job has been processed.
	JobStatusDone
	// JobStatusFailed means the job has failed to be processed.
	JobStatusFailed
)

// Job interface for job processing
type Job interface {
	Process()
	Status(status int, err error)
}

type simpleJob struct {
	task     func()
	callback func(status int, err error)
	status   int
	err      error
}

func (s *simpleJob) Process() {
	s.task()
}

func (s *simpleJob) Status(status int, err error) {
	s.status = status
	s.err = err

	s.callback(status, err)
}

// NewJob creates a new job
func NewJob(task func(), status ...func(status int, err error)) Job {
	statusX := func(status int, err error) {}
	if len(status) > 0 {
		statusX = status[0]
	}

	return &simpleJob{task: task, callback: statusX}
}
