package workerpool

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestPromiseError(t *testing.T) {
	type args struct {
		ctx    context.Context
		worker int
		limit  int
		tasks  int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "with 8 worker, 16 limits, 1000 task",
			args: args{
				ctx:    context.Background(),
				worker: 8,
				limit:  16,
				tasks:  10000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wp := NewAsync(tt.args.ctx, tt.args.worker, tt.args.limit)
			counter := atomic.Int32{}
			for i := 0; i < tt.args.tasks; i++ {
				promise := wp.Exec(context.Background(),
					func(ctx context.Context) error {
						parallelTasks := counter.Add(1)
						defer counter.Add(-1)

						if parallelTasks > int32(tt.args.limit) {
							t.Errorf("expected parallel tasks less than %d, but got %d", tt.args.limit, parallelTasks)
						}

						if parallelTasks < 1 {
							t.Errorf("expected parallel tasks greater than 0, but got %d", parallelTasks)
						}

						time.Sleep(time.Millisecond)
						return errors.New("something wrong")
					})
				go func() {
					assert.Error(t, promise.Error(), "expected error but got nil")
				}()
			}
			wp.Wait()
		})
	}
}

func TestWPAsyncExec(t *testing.T) {
	type args struct {
		ctx    context.Context
		worker int
		limit  int
		tasks  int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "with 8 worker, 16 limits, 1000 task",
			args: args{
				ctx:    context.Background(),
				worker: 8,
				limit:  16,
				tasks:  1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wp := NewAsync(tt.args.ctx, tt.args.worker, tt.args.limit)
			counter := atomic.Int32{}
			wg := sync.WaitGroup{}
			wg.Add(tt.args.tasks)
			for i := 0; i < tt.args.tasks; i++ {
				_ = wp.Exec(context.Background(),
					func(ctx context.Context) error {
						parallelTasks := counter.Add(1)
						defer counter.Add(-1)

						if parallelTasks > int32(tt.args.limit) {
							t.Errorf("expected parallel tasks less than %d, but got %d", tt.args.limit, parallelTasks)
						}

						if parallelTasks < 1 {
							t.Errorf("expected parallel tasks greater than 0, but got %d", parallelTasks)
						}

						time.Sleep(time.Millisecond)
						return nil
					})
			}
			wg.Wait()
		})
	}
}
