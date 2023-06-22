package workerpool

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestWPExec(t *testing.T) {
	type args struct {
		ctx      context.Context
		limit    int
		requests int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "checking counter",
			args: args{
				limit:    10,
				requests: 10002,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wp := NewSync(tt.args.limit)
			counter := atomic.Int32{}
			for i := 0; i < tt.args.requests; i++ {
				wp.Exec(context.Background(), func(ctx context.Context) {
					numberOfWorker := counter.Add(1)
					defer counter.Add(-1)

					if numberOfWorker > int32(tt.args.limit) {
						t.Errorf("expected number of worker less than %d, but got %d", tt.args.limit, numberOfWorker)
					}

					if numberOfWorker < 1 {
						t.Errorf("expected number of workers greater than 0, but got %d", numberOfWorker)
					}

					time.Sleep(time.Millisecond)
				})
			}
		})
	}
}
