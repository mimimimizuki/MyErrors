package myerror_test

import (
	// "strconv"
	"myerror"
	"testing"
	"errors"
	"strconv"
	"github.com/stretchr/testify/require"
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

	_, errcase1 := strconv.Atoi("ss")
	errcase2 := doError()
	if ok := errors.As(errcase2, &te); !ok {
		t.Errorf("doError()'s type is not same with testError")
	}
	cases := map[string] struct {
		errorcase error
		want bool
	}{
		"strconv error": {errcase1, false},
		"original error": {errcase2, true},
		"": {errors.New("test"), false},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T){
			err, ok := myerror.As[*testError](tt.errorcase)
			require.Equal(t, ok, tt.want)
			if !ok {
				require.Nil(t, err)
			}
		})
	}
}
