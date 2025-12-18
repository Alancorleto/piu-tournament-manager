package json

import (
	"encoding/json"
	"net/http"
)

func ParseRequestParameters[T any](r *http.Request) (T, error) {
	var params T
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		return params, err
	}
	return params, nil
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

	return nil
}

func RespondWithError(w http.ResponseWriter, code int, msg string) error {
	return RespondWithJSON(w, code, map[string]string{"error": msg})
}
