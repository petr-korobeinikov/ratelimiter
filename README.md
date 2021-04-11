# ratelimiter

> The dumbest rate limiter I ever wrote to pass Golang interview.

## Usage

```go
rl := ratelimiter.New(
    WithDuration(10*time.Millisecond),
    WithTasks(
        func () { ... },
        func () { ... },
        func () { ... },
    ),
)
rl.Execute(ctx)
```
