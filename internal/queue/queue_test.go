package queue

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hibiken/asynq"
)

func TestHandleExample(t *testing.T) {
	b, _ := json.Marshal(ExamplePayload{Message: "hi"})
	task := asynq.NewTask(TypeExample, b)
	if err := HandleExample(context.Background(), task); err != nil {
		t.Fatalf("handle error: %v", err)
	}
}

func TestHandleExampleBadPayload(t *testing.T) {
	task := asynq.NewTask(TypeExample, []byte("not-json"))
	if err := HandleExample(context.Background(), task); err == nil {
		t.Fatal("expected error")
	}
}