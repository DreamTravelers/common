package err_code

import "fmt"

type LazyCatError struct {
	errorNo  int64
	errorMsg string
}

// Is 判断是否为某个具体错误
func (e LazyCatError) Is(err LazyCatError) bool {
	return err.errorNo == e.errorNo
}

// Error 错误信息
func (e LazyCatError) Error() string {
	return fmt.Sprintf("err_no = %d err_msg = %s", e.errorNo, e.errorMsg)
}

// ErrorNo 错误码
func (e LazyCatError) ErrorNo() int64 {
	return e.errorNo
}

// ErrorMsg 错误描述
func (e LazyCatError) ErrorMsg() string {
	return e.errorMsg
}

// WithCustomErrMsg 自定义某个错误的错误描述
func (e LazyCatError) WithCustomErrMsg(errMsg string) LazyCatError {
	return LazyCatError{
		errorNo:  e.ErrorNo(),
		errorMsg: errMsg,
	}
}

// GetFromErrorNo 通过错误码获取错误
func GetFromErrorNo(errorNO int64) LazyCatError {
	if err, ok := errorMap[errorNO]; ok {
		return err
	}
	return UnDefinedError
}
