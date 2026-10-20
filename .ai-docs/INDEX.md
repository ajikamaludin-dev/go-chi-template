# .ai-docs — Documentation Index

Entry point for AI agents. **Start here**, then open matching doc.

Last verified: 2026-10-20.

## Quick routing

| I want to… | Read |
|---|---|
| Understand system | architecture/system-overview.md |
| Backend file map | backend/backend-map.md |
| Write code patterns | patterns/coding-patterns.md |
| Add feature step-by-step | workflows/new-feature-recipe.md |
| Debug | workflows/debugging-playbook.md |
| Write/run tests | testing/testing-strategy.md |
| Jobs/queue/scheduler | backend/jobs-and-scheduler.md |
| AI workflow | workflows/ai-development-workflow.md |

## Directory map

```
.ai-docs/
├── INDEX.md
├── architecture/
├── backend/
├── database/
├── patterns/
├── workflows/
├── testing/
```

## Conventions

- Layering: Handler → Service → Repository → DB.
- Config: cleanenv struct tags.
- Tests: `go test ./...`.