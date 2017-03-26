package main

import (
	"errors"
	"strings"
)

// StringService to uppercase and cound chars in a string.
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return s, ErrEmpty
	}

	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty returned when string is empty
var ErrEmpty = errors.New("Empty string")

type ServiceMiddleware func(StringService) StringService
