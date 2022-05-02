# JobQueue - Powerful unlimited job queue with goroutine pool

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/jobqueue)](https://pkg.go.dev/github.com/go-zoox/jobqueue)
[![Build Status](https://github.com/go-zoox/jobqueue/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/jobqueue/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/jobqueue)](https://goreportcard.com/report/github.com/go-zoox/jobqueue)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/jobqueue/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/jobqueue?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/jobqueue.svg)](https://github.com/go-zoox/jobqueue/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/jobqueue.svg?label=Release)](https://github.com/go-zoox/jobqueue/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/jobqueue
```

## Getting Started

```go
import (
  "testing"
  "github.com/go-zoox/jobqueue"
)

func main(t *testing.T) {
	q := New(runtime.NumCPU())
	q.Start()
	defer q.Stop()

	for i := 0; i < 10; i++ {
		q.AddJob(&NewJob(func() {
			fmt.Println("Process Job")
		}))
	}
}
```

## Inspired by
* [dirkaholic/kyoo](https://github.com/dirkaholic/kyoo) - Unlimited job queue for go, using a pool of concurrent workers processing the job queue entries

## Related
* [go-zoox/cocurrent](https://github.com/go-zoox/cocurrent) - A Simple Goroutine Limit Pool
* [go-zoox/waitgroup](https://github.com/go-zoox/waitgroup) - Parallel-Controlled WaitGroup
* [go-zoox/promise](https://github.com/go-zoox/promise) - JavaScript Promise Like with Goroutines

## License
GoZoox is released under the [MIT License](./LICENSE).
