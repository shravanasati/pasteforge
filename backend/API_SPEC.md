# API Specification

This document lays out the API documentation for the pasteforge API.

Base URL: `localhost:3122`.

## `/api`

All API calls are prefixed with the `/api` path. Other paths on the same level `/` will either respond with a static asset (from the dist folder) or a 404 not found.

### `/api/v1`

The current latest version of the API is v1.

#### GET `/ping`

Returns 

```json
{"message": "pong"}
```

Here's what the full request path would look like: `http://localhost:3122/api/v1/ping`.

#### POST `/new`

Upload a paste. Requires the following JSON payload.

```json
{
	"content": "the paste content",
	"settings": {
		"language": "look at the acceptable languages below",
		"expirationDuration": "look at the acceptable expiration durations below",
		"expirationNumber": 15,
		"visibility": "look at the acceptable visibility values below",
		"password": "some password"
	}
}
```

Acceptable values for language:
```
"plain"
"css"
"cpp"
"go"
"html"
"java"
"javascript"
"json"
"python"
"rust"
"markdown"
"php"
"xml"
"yaml"
```

Acceptable values for expiration duration:
```
"minutes"
"hours"
"days"
"months"
"years"
"never"
```

Expiration number must be an integer greater than zero for all expiration durations except `never`.

Acceptable values for visibility:
```
"public"
"unlisted"
"private"
```

The whole `settings` field is optional. In case it is not present, the following default values would be used.

```json
{
	"settings": {
		"language": "plain",
		"expirationDuration": "never",
		"expirationNumber": 0,
		"visibility": "public",
		// no password would be enforced 
	}
}
```