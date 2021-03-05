# ratelimiter

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
