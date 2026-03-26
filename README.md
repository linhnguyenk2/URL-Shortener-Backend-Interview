### Running Locally

1. Start a PostgreSQL instance on port 5432 (or modify the `DATABASE_URL` env variable).
2. Run database migration (automatically executed on app start).
3. `go run cmd/server/main.go`
4. The server will start at `http://localhost:8080`


## API Endpoints

### 1. Create a Short URL

```bash
curl -X POST http://localhost:8080/api/shortlinks \
  -H "Content-Type: application/json" \
  -d '{"original_url": "https://example.com"}'
```

**Response:**
```json
{
  "id": "abc1234",
  "short_url": "http://localhost:8080/shortlinks/abc1234"
}
```

### 2. Get Short URL Detail

```bash
curl -X GET http://localhost:8080/api/shortlinks/abc1234
```

**Response:**
```json
{
  "id": "abc1234",
  "original_url": "https://example.com",
  "created_at": "2023-01-01T00:00:00Z"
}
```

### 3. Redirect to Original URL

```bash
curl -v http://localhost:8080/shortlinks/abc1234
```
Expect an HTTP 302 redirect back to the `original_url`.
