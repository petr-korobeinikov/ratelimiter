package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pkorobeinikov/ratelimiter"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	rl := ratelimiter.New(
		ratelimiter.WithInterval(750*time.Millisecond),
		ratelimiter.WithTasks(
			func() {
				fmt.Println("ok")
			},
			func() {
				fmt.Println("will never executed")
			},
		),
	)

	rl.Execute(ctx)
}
