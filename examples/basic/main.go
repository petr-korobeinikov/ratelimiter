package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pkorobeinikov/ratelimiter"
)

func main() {
	ctx := context.Background()

	rl := ratelimiter.New(
		ratelimiter.WithInterval(750*time.Millisecond),
		ratelimiter.WithTasks(
			func() {
				fmt.Println("three...")
			},
			func() {
				fmt.Println("two..")
			},
			func() {
				fmt.Println("one.")
			},
			func() {
				fmt.Println("done!")
			},
		),
	)

	rl.Execute(ctx)
}
