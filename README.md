# Golang Error Handling

Error handling

## Endpoints

`POST /status`

Example request body:
```JSON
{
    "name": "Example",
    "email": "example@gmail.com"
}
```

Required fields: `name`, `email`

## Error format

If a request fails any validation, expect errors in the following format:

```JSON
{
    "message": "There was a problem with validation",
    "status": 500,
    "error": "Internal Server Error",
    "causes": [
        "Name is required",
        "Email is required"
    ],
    "timestamp": "2021-01-31T14:43:20.4841902-03:00",
    "path": "/status"
}
```

## Error handling

### HttpException

```GO
NewHttpException(message string, status int, causes []string)
```

Required fields: `message`, `status`, `causes`

---

### NewBadRequestError

```GO
NewBadRequestError(message string, err error)
```

Returns status `400`

---

### NewInternalServerError

```GO
NewInternalServerError(message string, err error)
```

Returns status `500`

---

### NewNotAcceptableError

```GO
NewNotAcceptableError(message string, err error)
```

Returns status `406`

---

### NewNotFoundError

```GO
NewNotFoundError(message string)
```

Returns status `404`

---

### NewTooManyRequestsError

```GO
NewTooManyRequestsError(message string, err error)
```

Returns status `429`

---

### NewUnauthorizedError

```GO
NewUnauthorizedError(message string, err error)
```

Returns status `401`
