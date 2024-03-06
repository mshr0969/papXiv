package domain

import "errors"

var (
	ErrInvalidArgument  = errors.New("invalid argument")
	ErrNonExistentPaper = errors.New("non existent paper")
)
