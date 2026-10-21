package service

import (
	"context"

	"github.com/aji/chi-go-template/internal/queue"
	"github.com/hibiken/asynq"
)

// JobService wraps queue enqueue operations for HTTP handlers.
type JobService struct {
	client *asynq.Client
}

// NewJobService creates a JobService.
func NewJobService(client *asynq.Client) *JobService {
	return &JobService{client: client}
}

// EnqueueExample pushes an example task to the worker queue.
func (s *JobService) EnqueueExample(ctx context.Context, message string) error {
	return queue.EnqueueExample(ctx, s.client, message)
}
