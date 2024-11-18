package services

import "errors"

var (
	ErrInvalidData = errors.New("invalid data")

	ErrGenerateCode = errors.New("failed to generate code")
)
