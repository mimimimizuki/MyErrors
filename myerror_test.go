package myerror_test

import (
	// "strconv"
	"myerror"
	"testing"
	"fmt"
	"errors"
)

type testError struct {
	Message string
	N       int
}

func (t *testError) Error() string {
	return t.Message
}

func doError() *testError {
	return &testError{
		Message: "original error occurred", N: 1,
	}
}

func TestMyErrorsAs(t *testing.T) {
	var te *testError

	// _, errcase1 := strconv.Atoi("ss")
	errcase2result := doError()
	if ok := errors.As(errcase2result, &te); !ok {
		t.Errorf("doError()'s type is not same with testError")
	}
	cases := map[string] struct {
		errorcase error
		want bool
		errorresult error
	}{
		// "strconv error": {errcase1, false, errcase1},
		"original error": {errcase2result, true, errcase2result},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T){
			err, ok := myerror.As[*testError](tt.errorcase)
			fmt.Println(err, ok)
			if ok != tt.want {
				t.Error(err)
			}
			if tt.errorresult != err {
				t.Error(err)
			}
		})
	}
}
