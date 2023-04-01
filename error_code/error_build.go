package err_code

import "fmt"

var errorMap map[int64]LazyCatError

func buildError(errorNo int64, errorMsg string) LazyCatError {
	if errorMap == nil {
		errorMap = make(map[int64]LazyCatError)
	}
	if _, ok := errorMap[errorNo]; ok {
		panic(fmt.Sprintf("duplicated error code definition errorNo = %v", errorNo))
	}
	err := LazyCatError{
		errorNo:  errorNo,
		errorMsg: errorMsg,
	}
	errorMap[errorNo] = err

	return err
}
