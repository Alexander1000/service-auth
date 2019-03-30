package json

import (
	"encoding/json"
	"net/http"
)

func Reply(w http.ResponseWriter, data interface{}, code int) (int, error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	body, err := json.Marshal(data)
	if err != nil {
		body = []byte(`{"error": "Unknown and unpredictable error with huge, massive and catastrophic consequences!"}`)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(code)
	}

	return w.Write(body)
}
