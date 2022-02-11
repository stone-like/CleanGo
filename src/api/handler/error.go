package handler

import (
	"log"
	"net/http"

	"github.com/stonelike/CleanGo/src/api/codes"
	"github.com/stonelike/CleanGo/src/myerrors"
)

type generalError interface {
	Code() string
}

type internalError interface {
	Internal() bool
}

func HttpErrorResponse(w http.ResponseWriter, err error) {
	// _, genok := errors.Cause(err).(generalError)

	// if genok {
	// 	fmt.Printf("general Error: %s", err.Error())
	// }

	// _, internalok := errors.Cause(err).(internalError)

	// if internalok {
	// 	fmt.Printf("internal Error: %s", err.Error())
	// }

	// fmt.Printf("undefined Error: %s", err.Error())

	w.WriteHeader(codes.ToHttpCode(err))

	log.Printf("%+v", err)

	w.Write([]byte(myerrors.UserInfo(err)))

}
