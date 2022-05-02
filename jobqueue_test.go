package jobqueue

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

type JobX func()

func (j JobX) Process() {
	j()
}

type TestJobX struct {
	ID string
}

// Process - test process function
func (t *TestJobX) Process() {
	if t.ID == "3" {
		panic("job panic")
	}

	fmt.Printf("Processing job '%s'\n", t.ID)
	time.Sleep(1 * time.Second)
}

func (t *TestJobX) Status(status int, err error) {
	s := ""
	switch status {
	case JobStatusPending:
		s = "Pending"
	case JobStatusRunning:
		s = "Running"
	case JobStatusDone:
		s = "Done"
	case JobStatusFailed:
		s = "Failed"
	}

	fmt.Printf("Job '%s' status: %s \n", t.ID, s)
	if err != nil {
		fmt.Println(err)
	}
}

func TestJobQueue(t *testing.T) {
	// fmt.Println("runtime.NumCPU():", runtime.NumCPU())
	q := New(runtime.NumCPU())
	q.Start()
	defer q.Stop()

	for i := 0; i < 10; i++ {
		// q.AddJob(&TestJobX{strconv.Itoa(i + 1)})
		id := i
		q.AddJob(NewJob(func() {
			fmt.Printf("Processing job: %d\n", id)
			time.Sleep(3 * time.Second)
		}))
	}
}
