# AGENTS.md — chi-go-template

> Hidden dirs (`.ai-docs/`, `.github/`) invisible to glob. Use `ls -la` / `find`.

> **Start here:** `.ai-docs/INDEX.md`.

## Stack & Repo

Go 1.24 + chi v5 / sqlx (mysql/postgres/sqlite) / asynq (Redis) / cleanenv. Repo: `chi-go-template`.

## Architecture

**Layering:** `Handler → Service → Repository → DB`

- Handlers thin (HTTP only, chi). Business logic in `internal/service/`. Repositories in `internal/repository/`. Models in `internal/model/`. External APIs → `internal/service/external/`.
- **Never** import Handler from Service/Repository. **Never** query DB in handlers.
- Queue jobs in `internal/queue/` (asynq tasks + handlers). Worker binary `cmd/worker`.

## Backend

- **Handlers:** `internal/http/handler/{domain}/`. Constructor-inject Service. Return JSON. Validate via request DTO + manual checks.
- **Services:** one per aggregate. Plain Go, no binding needed.
- **Repositories:** sqlx queries, context-aware.
- **Models:** plain structs with db tags. Status constants on model.
- **Migrations:** `migrations/` numbered `.up.sql`/`.down.sql` via golang-migrate.
- **Transactions:** wrap multi-step ops in Service/Repository as needed.
- **Config:** cleanenv struct tags in `internal/config/`.
- **Jobs:** asynq tasks in `internal/queue/`. Schedule via asynq periodic tasks.

## Code Style

- gofmt / goimports. No docblock on simple funcs. Only "why" comments.
- Errors: wrap with %w, return up. Use sentinel errors where needed.
- Logging: slog via observability.Channels. No fmt.Println in committed code.
- Forbidden: panic in handlers, ignoring errors, global state.

## Security

- `.env` has real credentials — never commit. Never log secrets.

## Forbidden

- Bypass Service layer. DB queries in handlers.
- New deps without approval.
- Commit to main/master directly.
- `fmt.Println`/debug prints in committed code.

## AI Workflow

**Plan-first:** analyze → plan in `task/CURRENT.md` (status `planning`). See `.ai-docs/workflows/ai-development-workflow.md`.

**YOLO Mode:** user says `yolo` → skip plan, execute directly.

**SAO:** user says `sao` → answer as short as possible.

**YON:** user says `yon`/`yesOrNo`/`yes Or No` → answer yes or no only.

**No production code in plans.** Pseudocode only for complex logic.
**After coding:** run `go test ./...` + `go vet ./...` + `gofmt -l .`; set CURRENT.md `status: review`, present summary, **stop and wait**. After approval: `status: done`, move to `task/done/`, update LOG.md.

## Operating rules

- Act in **plan mode** by default — no edits/commands unless explicitly approved.
- Questions → answer or ask, never edit.
- **No git write operations.**
- Don't run tests/builds unless explicitly asked.
- Shortest answer first. Match pace. Ask before going deep.
