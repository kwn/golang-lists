package errors

import "net/http"

const errorType string = "http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html"

type HttpError struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func (e *HttpError) BadRequest(d string) {
	e.Type = errorType
	e.Title = "Bad request"
	e.Status = http.StatusBadRequest
	e.Detail = d
}

func (e *HttpError) Unauthorized(d string) {
	e.Type = errorType
	e.Title = "Unauthorized"
	e.Status = http.StatusUnauthorized
	e.Detail = d
}

func (e *HttpError) Forbidden(d string) {
	e.Type = errorType
	e.Title = "Forbidden"
	e.Status = http.StatusForbidden
	e.Detail = d
}

func (e *HttpError) NotFound(d string) {
	e.Type = errorType
	e.Title = "Not found"
	e.Status = http.StatusNotFound
	e.Detail = d
}

func (e *HttpError) MethodNotAllowed(d string) {
	e.Type = errorType
	e.Title = "Method Not Allowed"
	e.Status = http.StatusMethodNotAllowed
	e.Detail = d
}

func (e *HttpError) Conflict(d string) {
	e.Type = errorType
	e.Title = "Conflict"
	e.Status = http.StatusConflict
	e.Detail = d
}

func NewHttpError(code int) HttpError {
	return NewHttpErrorWithDetails(code, "")
}

func NewHttpErrorWithDetails(code int, details string) HttpError {
	err := &HttpError{}

	switch code {
	case http.StatusBadRequest:
		err.BadRequest(details)
	case http.StatusUnauthorized:
		err.Unauthorized(details)
	case http.StatusForbidden:
		err.Forbidden(details)
	case http.StatusNotFound:
		err.NotFound(details)
	case http.StatusMethodNotAllowed:
		err.MethodNotAllowed(details)
	case http.StatusConflict:
		err.Conflict(details)
	}

	if details == "" {
		err.Detail = err.Title
	}

	return *err
}
