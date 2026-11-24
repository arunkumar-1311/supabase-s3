# supabase-s3

A small Go HTTP API that exposes S3-style operations (upload, delete, presigned download URLs) on top of [Supabase Storage](https://supabase.com/docs/guides/storage). Built with [Gin](https://github.com/gin-gonic/gin).

## Requirements

- Go (see `go.mod` for the version)
- A Supabase project with a storage bucket

## Configuration

Copy `.env.example` to `.env` and fill in the values:

| Variable          | Required | Description                                   |
| ----------------- | -------- | --------------------------------------------- |
| `SUPABASE_URL`    | yes      | Your project URL, e.g. `https://xyz.supabase.co` |
| `SUPABASE_KEY`    | yes      | Service-role or anon key                      |
| `SUPABASE_BUCKET` | yes      | Name of the storage bucket to use             |
| `APP_ADDR`        | no       | Listen address (default `:8080`)              |

Environment variables already set in the shell take effect even without a `.env` file.

## Run

```sh
go run ./cmd/app
```

## API

### Upload a file

`POST /files` — multipart form with:

- `file` (required): the file to upload
- `key` (optional): object path in the bucket; defaults to the uploaded filename

Existing objects with the same key are overwritten (upsert).

```sh
curl -F "file=@photo.png" -F "key=images/photo.png" http://localhost:8080/files
```

Response `201`:

```json
{ "key": "images/photo.png", "message": "..." }
```

### Get a presigned URL

`GET /presigned-url/*key?expires_in=<seconds>` — returns a time-limited download URL. `expires_in` is optional (default `300`) and must be a positive integer.

```sh
curl "http://localhost:8080/presigned-url/images/photo.png?expires_in=600"
```

Response `200`:

```json
{ "key": "images/photo.png", "url": "https://...", "expires_in": 600 }
```

### Delete a file

`DELETE /files/*key`

```sh
curl -X DELETE http://localhost:8080/files/images/photo.png
```

Response `200`:

```json
{ "key": "images/photo.png", "deleted": true }
```

### Errors

Errors are returned as `{ "error": "<message>" }` with `400` for bad input and `502` when Supabase Storage returns an error.

## Project layout

```
cmd/app/            entry point
internal/api/       Gin router and HTTP handlers
internal/config/    environment-based configuration
internal/storage/   Supabase Storage client (upload, download, list, delete, URLs, buckets)
```
