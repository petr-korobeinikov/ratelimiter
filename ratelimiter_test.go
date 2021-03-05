package ratelimiter

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRatelimiter_Execute(t *testing.T) {
	t.Run(`postitive`, func(t *testing.T) {
		var called1, called2, called3 bool
		ctx := context.Background()

		sut := New(
			WithDuration(10*time.Millisecond),
			WithTasks(func() { called1 = true }, func() { called2 = true }, func() { called3 = true }),
		)
		sut.Execute(ctx)

		assert.True(t, called1)
		assert.True(t, called2)
		assert.True(t, called3)
	})

	t.Run(`no tasks supplied`, func(t *testing.T) {
		ctx := context.Background()

		sut := New(
			WithDuration(50 * time.Millisecond),
		)
		sut.Execute(ctx)
	})

	t.Run(`no duration supplied`, func(t *testing.T) {
		ctx := context.Background()

		sut := New()
		sut.Execute(ctx)
	})

	t.Run(`timeout`, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()

		var called1, called2 bool

		sut := New(
			WithDuration(35*time.Millisecond),
			WithTasks(func() { called1 = true }, func() { called2 = true }),
		)
		sut.Execute(ctx)

		assert.True(t, called1)
		assert.False(t, called2)
	})
}
