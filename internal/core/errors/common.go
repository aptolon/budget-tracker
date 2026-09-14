package core_errors

import "errors"

var (
	// 400
	ErrInvalidArgument = errors.New("invalid argument")
	//401
	ErrUnauthorized = errors.New("unauthorized")
	//403
	ErrForbidden = errors.New("forbidden")
	//404
	ErrNotFound = errors.New("not found")
	//409
	ErrConflict = errors.New("conflict")
	//500
	ErrInternal = errors.New("internal")
)
