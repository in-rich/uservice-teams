package dao

import "errors"

var (
	ErrMemberAlreadyExists = errors.New("member already exists")
	ErrMemberNotFound      = errors.New("member not found")

	ErrTeamNotFound = errors.New("team not found")

	ErrCodeAlreadyExists = errors.New("code already exists")
	ErrCodeNotFound      = errors.New("code not found")
)
