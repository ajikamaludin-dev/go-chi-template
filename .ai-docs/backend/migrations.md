# Migrations

## Run in production

Migrations are embedded into the `migrate` binary (VFS), so only the binary is needed.

1. Ensure `.env` has correct `DB_*` settings.
2. Build binary:
   ```bash
   go build -o migrate ./cmd/migrate
   ```
3. Run:
   ```bash
   ./migrate up
   ```
4. Roll back if needed:
   ```bash
   ./migrate down
   ```

## Rules

- Migrations are versioned; never edit applied migrations in prod.
- Test migrations in staging first.
- Backup DB before migrating.
- Keep migrations idempotent where possible.