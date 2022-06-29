# URL Shortener

## To run

server is bound to port 8080

```
go run cmd/server/main.go
```

## Design decisions

- HTTP server is separated into its own package so that all http related things are isolated in its own package.
- Shortener is done with the memory store, but it can be refactored out in the future
- In package root, are all the shared types and sentient errors
- Shortener is a sequential shortener, it will just go sequentially as it is the simplest to implement for the
  constraints, and prevents any potential conflict

## API

### /shorten

is used to shorten the url

#### Request Format

- Method: POST
- Content-Type: application/json
- Body format

```
{
    "url": "http://www.example.com/long-url-path"
}
```

#### Response Format

- Content-Type: application/json
- Body format

```
{
    "url": "localhost:8080/A-----"
}
```

### /{key}

#### Request Format

- Method: GET
- {key}: the key returned by the shorten API

#### Response format

- N/A
- Redirected to url
- 404 if not found

### /stats/{key}

#### Request Format

- Method: GET
- {key}: the key returned by the shorten API

#### Response Format

- Body format

```
{
	"url": "https://www.google.com",
	"count": 3,
	"createdAt": "2022-06-29T16:07:20.545742+08:00"
}
```
