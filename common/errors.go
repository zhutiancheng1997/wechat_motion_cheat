package common

import (
	"fmt"
)

// ErrorCode error code
type ErrorCode struct {
	Code    int32  `json:"code"`
	Msg     string `json:"msg"`
	TraceID string `json:"trace_id,omitempty"`
}

func NewErrorCodeByET(err error, traceID string) ErrorCode {
	return ErrorCode{-1, err.Error(), traceID}
}

func NewErrorCodeByError(err error) ErrorCode { return ErrorCode{-1, err.Error(), ""} }

func NewErrorCodeByMsg(msg string) ErrorCode { return ErrorCode{-1, msg, ""} }

// NewErrorCode new ErrorCode
func NewErrorCode(code int32, msg string) ErrorCode { return ErrorCode{code, msg, ""} }

func (e ErrorCode) String() string {
	return fmt.Sprintf("code: %d,msg: %s", e.Code, e.Msg)
}
func (e ErrorCode) Error() string {
	return fmt.Sprintf("code: %d,msg: %s", e.Code, e.Msg)
}

// GetErrMsg get error message
func (e ErrorCode) GetErrMsg() string {
	return e.Msg
}

// GetErrCode get error code
func (e ErrorCode) GetErrCode() int32 {
	return e.Code
}
