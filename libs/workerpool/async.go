package workerpool

import (
	"context"
	"sync"
)

// WPAsync is a worker pool with async execution
// It contains a channel of tasks and a limit of pool and number of workers
type WPAsync struct {
	taskCh chan Task
	worker int
	limit  int
	wg     sync.WaitGroup
}

// Task is a task to be executed by a worker
// It contains the context of the task and the function to be executed
// It also contains a promise to get the result of the task
type Task struct {
	ctx     context.Context
	promise *Promise
	fn      func(ctx context.Context) error
}

// Promise is a promise of a task
// It is used to get the result of a task
type Promise struct {
	ch chan error
}

// NewAsync creates a new async worker pool
func NewAsync(ctx context.Context, worker int, limit int) *WPAsync {
	w := WPAsync{
		worker: worker,
		limit:  limit,
		taskCh: make(chan Task, limit),
	}

	go func() {
		<-ctx.Done()
		close(w.taskCh)
	}()

	for i := 0; i < w.worker; i++ {
		go func() {
			for task := range w.taskCh {
				select {
				case <-task.ctx.Done():
					task.promise.ch <- task.ctx.Err()
				default:
					task.promise.ch <- task.fn(task.ctx)
				}
				w.wg.Done()
			}
		}()
	}
	return &w
}

// Exec add to the task channel a task to be executed by a worker
// It returns a promise to get the result of the task
func (w *WPAsync) Exec(ctx context.Context, task func(ctx context.Context) error) *Promise {
	w.wg.Add(1)
	promise := &Promise{
		ch: make(chan error, 1),
	}

	select {
	case <-ctx.Done():
		promise.ch <- ctx.Err()
		w.wg.Done()
	default:
		w.taskCh <- Task{
			ctx:     ctx,
			promise: promise,
			fn:      task,
		}
	}

	return promise
}

// Wait waits for all tasks to be done
func (w *WPAsync) Wait() {
	w.wg.Wait()
}

// Error returns the error of the task
// It closes the channel after return fetch error
func (p *Promise) Error() error {
	defer close(p.ch)
	return <-p.ch
}
