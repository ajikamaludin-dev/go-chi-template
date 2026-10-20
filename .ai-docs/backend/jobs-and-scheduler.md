# Jobs & Scheduler (asynq)

Queue backend: **asynq** (Redis). All job code lives in `internal/queue/`. Two binaries share the code: `cmd/web` (enqueues) and `cmd/worker` (processes).

## Concepts

- **Task type**: string constant, e.g. `example:task`. Defines a job kind.
- **Payload**: JSON-serializable struct.
- **Client**: `asynq.Client` used by web to enqueue.
- **Server**: `asynq.Server` running in worker, consumes tasks.
- **ServeMux**: maps task types to handler functions.
- **Scheduler**: `asynq.Scheduler` for periodic/cron jobs (runs in worker).

## How it works

1. Web calls `queue.EnqueueX(ctx, client, payload)` → task pushed to Redis.
2. Worker's `asynq.Server` pulls task, dispatches to registered handler via `ServeMux`.
3. Handler returns `nil` on success; non-nil error triggers retry with backoff.
4. After max retries, task moves to dead-letter (archived) queue.

## Adding a new job (rules)

1. **Define task type + payload** in `internal/queue/<domain>.go`:
   ```go
   const TypeSendEmail = "email:send"
   type SendEmailPayload struct {
       To      string `json:"to"`
       Subject string `json:"subject"`
       Body    string `json:"body"`
   }
   ```
2. **Enqueue helper**:
   ```go
   func EnqueueSendEmail(ctx context.Context, client *asynq.Client, p SendEmailPayload) error {
       b, _ := json.Marshal(p)
       t := asynq.NewTask(TypeSendEmail, b, asynq.MaxRetry(5))
       _, err := client.Enqueue(t)
       return err
   }
   ```
3. **Handler** (must be idempotent where possible):
   ```go
   func HandleSendEmail(ctx context.Context, t *asynq.Task) error {
       var p SendEmailPayload
       if err := json.Unmarshal(t.Payload(), &p); err != nil {
           return fmt.Errorf("unmarshal: %w", err)
       }
       // do work; return err to retry
       return nil
   }
   ```
4. **Register handler** in `queue.NewMux()`:
   ```go
   mux.HandleFunc(TypeSendEmail, HandleSendEmail)
   ```
5. **Wire client in web** (pass `*asynq.Client` to services that need enqueueing).
6. **Add tests** for handler logic (unit) and enqueue (functional).

## Periodic / scheduled jobs

Use `asynq.Scheduler` in worker:
```go
sched := asynq.NewScheduler(redisOpt, &asynq.SchedulerOpts{})
sched.Register("@every 1h", asynq.NewTask(TypeCleanup, nil))
sched.Run()
```
Run scheduler in a goroutine alongside server in `cmd/worker/main.go`.

## Retry & backoff

- Default max retry: 25. Override per task via `asynq.MaxRetry(n)`.
- Backoff: exponential by default. Custom via `asynq.RetryDelayFunc`.
- Idempotency: handlers should tolerate re-execution (use unique keys / DB upserts).

## Dead-letter / inspection

- Failed tasks beyond retries go to archived queue.
- Inspect with asynq CLI: `asynq dash` or `redis-cli`.
- Add monitoring/alerts on queue depth and failure rate (Prometheus metrics).

## Rules

- Task types are constants; never hardcode strings in multiple places.
- Payloads must be JSON-serializable; avoid pointers to DB objects.
- Handlers must not panic; recover and return error.
- Long jobs should be chunked or use context cancellation.
- Keep handlers in `internal/queue/`, business logic in services (call service from handler).
- Do not enqueue from worker to avoid loops unless intentional.
- Test handlers with unit tests; integration tests spin Redis via testcontainers.

## Where to add

- New domain file: `internal/queue/<domain>.go`
- Registration: `internal/queue/queue.go` → `NewMux()`
- Worker wiring: `cmd/worker/main.go`
- Client usage: inject `*asynq.Client` into service.