package queue

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

const TypeExample = "example:task"

type ExamplePayload struct {
	Message string `json:"message"`
}

func NewClient(redisAddr, password string, db int) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr, Password: password, DB: db})
}

func NewServer(redisAddr, password string, db int, concurrency int) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr, Password: password, DB: db},
		asynq.Config{Concurrency: concurrency},
	)
}

func EnqueueExample(ctx context.Context, client *asynq.Client, msg string) error {
	b, _ := json.Marshal(ExamplePayload{Message: msg})
	t := asynq.NewTask(TypeExample, b)
	_, err := client.Enqueue(t)
	return err
}

func HandleExample(ctx context.Context, t *asynq.Task) error {
	var p ExamplePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	fmt.Printf("processed example task: %s\n", p.Message)
	return nil
}

func NewMux() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeExample, HandleExample)
	return mux
}
