package workerpool

import (
	"context"
)

// WP is a worker pool with sync execution
type WP chan struct{}

// NewSync creates a new sync worker pool with a pool size
func NewSync(poolSize int) WP {
	return make(chan struct{}, poolSize)
}

// Exec executes a function in a worker
func (w WP) Exec(ctx context.Context, fn func(ctx context.Context)) {
	select {
	case <-ctx.Done():
	default:
		w.acquire()
		go func() {
			defer w.release()
			fn(ctx)
		}()
	}
}

// acquire acquires a worker
func (w WP) acquire() {
	w <- struct{}{}
}

// release releases a worker
func (w WP) release() {
	<-w
}
