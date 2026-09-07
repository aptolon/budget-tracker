package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/aptolon/budget-tracker/internal/core/errors"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	queryValues := r.URL.Query()
	paramValue := queryValues.Get(key)
	if paramValue == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(paramValue)
	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key='%s' not a valid integer: %v: %w",
			paramValue,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return &val, nil
}
