# CRUD Pattern (example)

## Layers

- Model: `internal/model/user.go` (struct with db tags)
- Repository: `internal/repository/user_repo.go` (sqlx queries)
- Service: `internal/service/user_service.go` (business logic)
- Handler: `internal/http/handler/user/handler.go` (chi)

## Repository example

```go
type UserRepo struct { db *sqlx.DB }
func (r *UserRepo) Get(ctx context.Context, id int64) (*model.User, error) {
    var u model.User
    err := r.db.GetContext(ctx, &u, "SELECT * FROM users WHERE id=?", id)
    return &u, err
}
```

## Service example

```go
type UserService struct { repo *repository.UserRepo }
func (s *UserService) Get(ctx context.Context, id int64) (*model.User, error) {
    return s.repo.Get(ctx, id)
}
```

## Handler example

```go
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
    u, err := h.svc.Get(r.Context(), id)
    if err != nil { http.Error(w, err.Error(), 500); return }
    json.NewEncoder(w).Encode(u)
}
```

## Route registration

```go
r.Route("/api/v1/users", func(r chi.Router) {
    r.Get("/{id}", handler.Get)
    r.Post("/", handler.Create)
})
```