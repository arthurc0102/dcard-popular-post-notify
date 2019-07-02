package helper

import (
	"errors"
	"strings"
)

// MergeErrors merge error array to one error
func MergeErrors(errs []error) error {
	var message []string

	for _, err := range errs {
		message = append(message, err.Error())
	}

	return errors.New(strings.Join(message, ", "))
}
