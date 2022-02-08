package handler

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type generalError interface {
	Code() string
}

type internalError interface {
	Internal() bool
}

func ErrorResponse(w http.ResponseWriter, err error) {
	_, genok := errors.Cause(err).(generalError)

	if genok {
		fmt.Printf("general Error: %s", err.Error())
	}

	_, internalok := errors.Cause(err).(internalError)

	if internalok {
		fmt.Printf("general Error: %s", err.Error())
	}

	fmt.Printf("undefined Error: %s", err.Error())

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))

}
