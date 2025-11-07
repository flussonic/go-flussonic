package apierror

import (
	"encoding/json"
	"fmt"
)

// Description: an object containing references to the source of the error
type ErrorSource struct {
	// Description: a string indicating which URI query parameter caused the error.
	Parameter *string `json:"parameter,omitempty" validate:"omitempty"`
	// Description: a JSON Pointer [RFC6901] to the associated entity in the request document
	// [e.g. `"/data"` for a primary data object, or `"/data/attributes/title"` for a specific attribute].
	Pointer *string `json:"pointer,omitempty" validate:"omitempty"`
}

// Error is a custom error structure returned by services.
type Error struct {
	ID     *string           `json:"id,omitempty"`
	Meta   map[string]string `json:"meta,omitempty"`
	Source *ErrorSource      `json:"source,omitempty"`
	Status *string           `json:"status,omitempty"`
	Title  *string           `json:"title,omitempty"`
	Code   *string           `json:"code,omitempty"`
}

// Error makes the structure compatible with the standard error interface.
func (e *Error) Error() string {
	if e.Title != nil && e.Code != nil {
		return fmt.Sprintf("API Error: %s (code: %s, status: %s)", *e.Title, *e.Code, *e.Status)
	}
	if e.Title != nil {
		return fmt.Sprintf("API Error: %s", *e.Title)
	}
	// Fallback if error body was empty
	return "An unknown API error occurred"
}

// NewFromBytes creates an Error instance from a byte slice (response body).
func NewFromBytes(body []byte) error {
	var apiErr Error
	if err := json.Unmarshal(body, &apiErr); err != nil {
		// If failed to parse error body, return it as plain text.
		return fmt.Errorf("failed to parse api error: %s", string(body))
	}
	return &apiErr
}

type ErrorResponse struct {
	Errors []Error `json:"errors,omitempty"`
}

// Error makes ErrorResponse compatible with the error interface.
// Returns combined messages from all errors in the array.
func (e *ErrorResponse) Error() string {
	if len(e.Errors) == 0 {
		return "An unknown API error occurred"
	}

	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}

	// If multiple errors, combine them
	result := fmt.Sprintf("Multiple API errors (%d): ", len(e.Errors))
	for i, err := range e.Errors {
		if i > 0 {
			result += "; "
		}
		result += err.Error()
	}
	return result
}
