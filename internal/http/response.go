package http

import (
	"context"
	"encoding/json"
	"net/http"
)

func (s *server) writeSuccess(ctx context.Context, w http.ResponseWriter, data interface{}, status int) error {
	if status != 0 {
		w.WriteHeader(status)
	}

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *server) writeError(ctx context.Context, w http.ResponseWriter, err error, code int) error {
	if http.StatusText(code) == "" {
		return ErrUnrecognizedCode
	}

	if err == nil {
		err = ErrUnknownError
	}

	s.logger.Errorf("request failed: %v", err.Error())
	w.WriteHeader(code)

	res := struct {
		Message string `json:"message"`
	}{
		Message: http.StatusText(code),
	}

	return json.NewEncoder(w).Encode(res)
}
