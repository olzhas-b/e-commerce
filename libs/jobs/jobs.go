package jobs

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type Job interface {
	Start(ctx context.Context, dur time.Duration)
	Stop()
	Error() <-chan error
}

type job struct {
	stopSignal chan struct{}
	errorChan  chan error
	mu         sync.Mutex
	wg         sync.WaitGroup
	fn         func(ctx context.Context) error
	isStarted  atomic.Int32
	trackError bool
}

const (
	_defaultTrackError = false
)

const (
	statusStopped = iota
	statusStarted
)

func New(fn func(context.Context) error, opts ...Option) Job {
	j := &job{
		fn:         fn,
		stopSignal: make(chan struct{}, 1),
		errorChan:  make(chan error),
		trackError: _defaultTrackError,
	}

	for _, opt := range opts {
		opt(j)
	}

	return j
}

func (j *job) Start(ctx context.Context, dur time.Duration) {
	go func() {
		j.mu.Lock()
		if j.isStarted.Load() == statusStarted {
			j.mu.Unlock()
			return
		}
		j.isStarted.Store(statusStarted)
		defer j.isStarted.Store(statusStopped)
		j.mu.Unlock()

		ticker := time.NewTicker(dur)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				j.handleError(j.fn(ctx))
			case <-j.stopSignal:
				return
			}
		}
	}()
}

func (j *job) handleError(err error) {
	if err != nil && j.trackError {
		j.errorChan <- err
	}
}

func (j *job) Stop() {
	j.stopSignal <- struct{}{}
}

func (j *job) Error() <-chan error {
	return j.errorChan
}
