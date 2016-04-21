package main

import (
	"errors"
	"strings"
)

// Validater interface
type Validater interface {
	// The Validate method showed by Validater is the core of this library
	Validate() error
}

// A set of several validaters
type ValidaterSet struct {
	validaters []Validater
}

// Returns a new validater set from the one given in parameters
func NewValidaterSet(v ...Validater) *ValidaterSet {
	set := make([]Validater, 0)
	set = append(set, v...)
	return &ValidaterSet{set}
}

// ValidaterSet returns a chain of all the error it encountered if any, or nil if it hasn't
func (v *ValidaterSet) Validate() error {
	acc := ""
	for _, valid := range v.validaters {
		if err := valid.Validate(); err != nil {
			acc += err.Error() + "\n"
		}
	}
	if acc == "" {
		return nil
	}
	return errors.New(strings.TrimRight(acc, "\n"))
}

// Exec returns the first error encountered or nil if it has not
func Exec(v ...Validater) error {
	for _, valid := range v {
		if err := valid.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Collect collects all of the errors and returns them, nil if no error were encountered
func Collect(v ...Validater) error {
	return (&ValidaterSet{v}).Validate()
}
