# Coding Patterns

- Handler: parse request, call service, write JSON response.
- Service: orchestrate repositories, transactions.
- Repository: sqlx queries, context-aware.
- Errors: return wrapped errors; handlers map to HTTP status.
- Logging: use observability channels (App/Error/Audit).