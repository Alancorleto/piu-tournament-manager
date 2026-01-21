package server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func mustPathUUID(r *http.Request, key string) (uuid.UUID, error) {
	value := r.PathValue(key)
	if value == "" {
		return uuid.Nil, fmt.Errorf("missing %s", key)
	}

	id, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid %s format", key)
	}

	return id, nil
}
