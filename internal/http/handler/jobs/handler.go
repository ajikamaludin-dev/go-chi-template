package jobs

import (
	"encoding/json"
	"net/http"

	"github.com/aji/chi-go-template/internal/service"
)

// Handler exposes HTTP endpoints for enqueueing jobs.
type Handler struct {
	svc *service.JobService
}

// New creates a jobs handler.
func New(svc *service.JobService) *Handler {
	return &Handler{svc: svc}
}

type enqueueExampleRequest struct {
	Message string `json:"message"`
}

// EnqueueExample handles POST /api/v1/jobs/example.
func (h *Handler) EnqueueExample(w http.ResponseWriter, r *http.Request) {
	var req enqueueExampleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.Message == "" {
		http.Error(w, "message is required", http.StatusBadRequest)
		return
	}

	if err := h.svc.EnqueueExample(r.Context(), req.Message); err != nil {
		http.Error(w, "failed to enqueue job", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "queued"})
}
