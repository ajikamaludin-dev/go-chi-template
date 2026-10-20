# AI Development Workflow

## Plan-First Rule (Mandatory)

For any task that changes code, AI agents MUST NOT write code before the plan is approved.

1. Analyze — read AGENTS.md, relevant .ai-docs, inspect neighboring files, search before creating, analyze impact.
2. Plan — write into `task/CURRENT.md` with `status: planning`.
3. Present — summarize plan, ask approval.
4. Wait — no edits until approved.
5. Approve — set `status: plan_approved` then `status: in_progress`.

Exceptions: trivial one-line fixes, typos, or user says "just do it"/"skip the plan".

## Roles

| Task type | Role | Focus |
|---|---|---|
| Backend (handler/service/repo/migration/job) | Senior Backend Engineer | service layer, validation, DB schema, tests |
| Infra/deploy/release | DevOps Engineer | migrations, docker, changelog |

## Verification (After Coding)

1. Run tests: `go test ./...`
2. Run vet: `go vet ./...`
3. Format: `gofmt -l .` (must be empty)
4. Self-review: no business logic in handlers, no ignored errors.
5. Acceptance criteria all checked.

## Task Tracking

`task/CURRENT.md` active. On completion: set `status: done`, move to `task/done/YYYY-MM-DD - Title.md`, append LOG.md and done/INDEX.md.
