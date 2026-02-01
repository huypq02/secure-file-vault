# Secure File Storage with S3 and PostgreSQL

A Go project for managing encrypted file uploads and downloads. Built with clean architecture principles because I got tired of refactoring monoliths. Uses AWS S3 for storage and PostgreSQL for metadata.

## What It Does

**Upload files**: You upload a file, it gets encrypted with AES-256, and stored in S3. Metadata goes into Postgres.

**Download files**: Fetch from S3, decrypt, and serve back to the user. Works, doesn't explode.

**Expiring shares**: Generate time-limited shareable links. Good for sending sensitive docs without them lingering forever.

**Auto cleanup**: Cron job deletes expired files from S3 and the database. Still figuring out the best way to handle edge cases here (concurrent deletes are annoying).

**Auth**: JWT tokens for now. Nothing fancy, but it works.

---

## Architecture

Split into layers because it makes testing easier:

- **Domain**: Business logic and interfaces. Doesn't know about HTTP or S3.
- **Usecase**: Application business logic. Orchestrates domain and infrastructure.
- **Interface**: HTTP handlers and routing.
- **Infrastructure**: S3, database, crypto stuff. Anything that talks to external systems.

It's not perfect, but it's maintainable. And testable. Which matters more than you'd think.

## Setup

See [docs/SETUP.md](docs/SETUP.md). Pretty straightforward if you have AWS and Postgres already.

## Security Notes

- Files are encrypted before leaving the server
- Keys are... well, they're in environment variables for now. Not ideal for production but it works
- Access control is enforced at the handler level
- Expiry cleanup runs periodically to avoid accumulating old stuff

## Known Issues / TODO

- Key rotation isn't implemented yet. Something to think about.
- Concurrent delete operations on S3 and DB need better handling. Had some weird race conditions during testing.
- Multipart upload for large files isn't done. File size limits exist for now.
- Admin endpoints for user/file management are missing.
- Error messages could be more helpful.

## Testing

Basic test coverage exists. Not comprehensive. Some database tests use mocks, some use real Postgres in tests. It's a mess but it works.

Run with: `go test ./...`

## Running It

```bash
go run ./cmd/main.go
```

Server starts on port 8080 (or whatever you set in config). Hits the database and S3, hopefully succeeds.

---

## Why Clean Architecture?

Honestly? I got tired of fat controllers and database logic scattered everywhere. This structure makes it easier to:

- Test business logic without mocking the entire world
- Swap out S3 for something else if needed (though you probably won't)
- Add new features without breaking everything
- Onboard someone else onto the codebase

Is it overkill for a small project? Maybe. But it's nice.

---

## Lessons Learned

- Go's error handling is verbose but it forces you to actually think about failures
- S3's eventual consistency is annoying when you need immediate deletes
- JWT tokens in headers are simpler than cookies, but you still need HTTPS
- Database migrations should be automated (still working on this)

---

Feel free to open issues or PRs if you find problems or think something's stupid.
