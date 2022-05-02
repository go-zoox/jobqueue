package jobqueue

import "testing"

func TestNewJob(t *testing.T) {
	processed := false

	job := NewJob(func() {
		processed = true
	}, func(status int, err error) {
		t.Logf("Job status: %d, err: %v", status, err)
	})

	job.Process()

	if !processed {
		t.Error("task should be done")
	}

}
