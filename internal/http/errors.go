package http

import "errors"

var (
	ErrUnknownError     = errors.New("unknown error occurred")
	ErrUnrecognizedCode = errors.New("unrecognized status code")
)
