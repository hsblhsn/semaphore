package semaphore

import (
	"context"
	"golang.org/x/sync/semaphore"
	"runtime"
)

// Semaphore is an implementation of weighted semaphore.
// Uses `golang.org/x/sync/semaphore` under the hood.
type Semaphore struct {
	max       int64
	semaphore *semaphore.Weighted
}

// New returns a new Semaphore with the given max weight.
func New(max int64) *Semaphore {
	return &Semaphore{
		max:       max,
		semaphore: semaphore.NewWeighted(max),
	}
}

// Add acquires the semaphore with a weight of n
func (s *Semaphore) Add(ctx context.Context, n int64) error {
	return s.semaphore.Acquire(ctx, n)
}

// Done releases the semaphore with a weight of 1.
func (s *Semaphore) Done() {
	s.semaphore.Release(1)
}

// DoneN releases the semaphore with a weight of n.
func (s *Semaphore) DoneN(n int64) {
	s.semaphore.Release(n)
}

// Wait for the semaphore to be released.
func (s *Semaphore) Wait(ctx context.Context) error {
	return s.semaphore.Acquire(ctx, s.max)
}

// Exit calls s.Done() first then calls runtime.Goexit().
// If called inside a goroutine, the goroutine will exit immediately.
// See https://golang.org/pkg/runtime/#Goexit for the documentation.
func (s *Semaphore) Exit() {
	s.Done()
	runtime.Goexit()
}
