package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	// - json.NewDecoder(r.Body) creates a JSON decoder that reads from the request body stream (r.Body).
	//   It allows reading large or streaming JSON data without loading the entire body into memory at once.
	// - Decode(payload) reads the JSON from the stream and unmarshals it into the Go variable passed as payload.
	//   The payload must be a pointer to a struct, map, or any Go data structure that matches the JSON format.
	// If the JSON is invalid or does not match the structure of the payload, Decode returns an error.
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// - w.Header().Add("Content-Type", "application/json") sets the Content-Type header to indicate a JSON response.
	// - w.WriteHeader(status) sets the HTTP status code for the response.
	// - json.NewEncoder(w).Encode(v) converts the Go value v to JSON and writes it to the response body.
	// If encoding fails, the function returns an error.
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	// the map syntax part means that the response JSON would be
	// {
	// 		"error": "error message"
	// }
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
