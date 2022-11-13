package semaphore_test

import (
	"context"
	"github.com/hsblhsn/semaphore"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	t.Parallel()

	var (
		sem     = semaphore.New(10)
		counter int32
	)

	for range make([]struct{}, 10) {
		sem.Add(context.TODO(), 1)

		go func() {
			defer sem.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	sem.Wait(context.TODO())
	require.EqualValues(t, 10, atomic.LoadInt32(&counter))

	go func() {
		sem.Exit()
		atomic.AddInt32(&counter, 100)
	}()

	// The counter did not increment because the sem was exited before the increment.
	require.EqualValues(t, 10, atomic.LoadInt32(&counter))
}
