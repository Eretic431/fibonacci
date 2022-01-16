package models

import "errors"

var (
	ErrInvalidArguments = errors.New("invalid arguments")

	ErrNoRecord = errors.New("found no record")
)
