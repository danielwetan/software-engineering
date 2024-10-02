package model

import (
	"errors"
	"fmt"
)

var (
	ErrFailedParseRequestBody   = errors.New("failed to parse request body")
	ErrFailedToPrepareStatement = errors.New("failed to prepare statement")
	ErrDataNotFound             = errors.New("data not found")
)

// ErrorResponse is an error response
// swagger:model
type ErrorResponse struct {
	// The error message
	// required: true
	Error string `json:"error"`
}

// ErrNotFound is an error type that can be returned by store APIs
// when a query unexpectedly fetches no records.
type ErrNotFound struct {
	entity string
}

// NewErrNotFound creates a new ErrNotFound instance.
func NewErrNotFound(entity string) *ErrNotFound {
	return &ErrNotFound{
		entity: entity,
	}
}

func (nf *ErrNotFound) Error() string {
	return fmt.Sprintf("%s not found", nf.entity)
}

// IsErrNotFound returns true if `err` is or wraps one of:
// - model.ErrNotFound
func IsErrNotFound(err error) bool {
	if err == nil {
		return false
	}

	// check if this is a model.ErrNotFound
	var nf *ErrNotFound
	return errors.As(err, &nf)
}
