package ratelimiter

import (
	"context"
	"time"
)

func (r *ratelimiter) Execute(ctx context.Context) {
	if len(r.tasks) == 0 {
		return
	}

	first, tail := r.tasks[0], r.tasks[1:]

	first()

	for _, task := range tail {
		select {
		case <-ctx.Done():
			return
		case <-r.ticker.C:
			task()
		}
	}
}

func New(opts ...Option) *ratelimiter {
	r := &ratelimiter{
		ticker: time.NewTicker(1 * time.Millisecond),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func WithInterval(d time.Duration) Option {
	return func(r *ratelimiter) {
		r.ticker = time.NewTicker(d)
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
	ticker *time.Ticker
	tasks  []Task
}
