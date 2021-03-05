package ratelimiter

import (
	"context"
	"time"
)

func (r *ratelimiter) Execute(ctx context.Context) {
	for _, task := range r.tasks {
		select {
		case <-ctx.Done():
			return
		case <-r.timeCh:
			<-r.timeCh
			task()
		}
	}
}

func New(opts ...Option) *ratelimiter {
	r := &ratelimiter{
		timeCh: time.Tick(1 * time.Millisecond),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func WithInterval(d time.Duration) Option {
	return func(r *ratelimiter) {
		r.timeCh = time.Tick(d)
	}
}

func WithTasks(tasks ...Task) Option {
	return func(r *ratelimiter) {
		r.tasks = tasks
	}
}

type Option func(*ratelimiter)
type Task func()

type ratelimiter struct {
	timeCh <-chan time.Time
	tasks  []Task
}
