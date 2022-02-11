package myerrors

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/stonelike/CleanGo/src/codes"
)

type myError struct {
	code     codes.Code
	userInfo string
	err      error
}

func (e myError) Error() string {
	return fmt.Sprintf("Code: %s,Msg: %s", e.code, e.err)
}

func Errorf(c codes.Code, format string, values ...interface{}) error {

	//OKの場合はエラーじゃないので
	if c == codes.OK {
		return nil
	}

	return myError{
		code:     c,
		userInfo: fmt.Sprintf(format, values...),
		err:      errors.Errorf(format, values...),
	}
}

func UserInfo(err error) string {
	if err == nil {
		return ""
	}

	var e myError
	if errors.As(err, &e) {
		return e.userInfo
	}

	//myError以外はエラーとして使わない想定(myErrorをWrapして、myErrorをベースとして使う)
	return "Undefined Error"

}

//エラーからコードを抽出する、myErrorの場合抽出、その他のエラーの場合はUnknown
func Code(err error) codes.Code {
	if err == nil {
		return codes.OK
	}

	var e myError
	if errors.As(err, &e) {
		return e.code
	}

	return codes.Unknown
}

//スタックトレース
func StackTrace(err error) string {
	var e myError
	if errors.As(err, &e) {
		return fmt.Sprintf("%v\n", e.err)
	}

	return ""
}
