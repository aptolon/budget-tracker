package core_errors

import "errors"

var (
	// 400
	ErrInvalidArgument = errors.New("invalid argument")
	//401
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	//403
	ErrForbidden = errors.New("forbidden")
	//404
	ErrNotFound = errors.New("not found")
	//409
	ErrConflict   = errors.New("conflict")
	ErrLoginTaken = errors.New("login already taken")

	//500
	ErrInternal = errors.New("internal")
)
