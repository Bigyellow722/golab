package comic

import (
	"fmt"
	"net/http"
)

type StatusError struct {
	StatusCode int
	StatusText string
}

func newStatusError(code int) StatusError {
	return StatusError{
		StatusCode: code,
		StatusText: http.StatusText(code),
	}
}

//golang中的error类型是一个接口，主要是要实现Error()方法
func (e StatusError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.StatusText)
}
