package codes

import (
	"net/http"

	"github.com/stonelike/CleanGo/src/codes"
	"github.com/stonelike/CleanGo/src/myerrors"
)

func ToHttpCode(err error) int {

	var httpStatus int

	c := myerrors.Code(err)
	switch c {
	case codes.OK:
		httpStatus = http.StatusOK
	case codes.InvalidRequest:
		httpStatus = http.StatusBadRequest
	case codes.NotFound:
		httpStatus = http.StatusNotFound
	case codes.Database, codes.Internal, codes.Unknown:
		httpStatus = http.StatusInternalServerError
	default:
		httpStatus = http.StatusInternalServerError

	}
	return httpStatus
}
