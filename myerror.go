package myerror

import (
	"errors"
)

func As[T any](e error) (T, bool) {
	var myerr T
	ok := errors.As(e, &myerr)
	return myerr, ok
}
