package common

var (
	// 0:ok
	Ok = ErrorCode{Code: 0, Msg: "success"}

	// [100, 999] common
	ErrorUnknown    = ErrorCode{Code: 100, Msg: "unknown error"}
	ErrorParams     = ErrorCode{Code: 101, Msg: "params invalid"}
	ErrorPermission = ErrorCode{Code: 102, Msg: "no permission"}
	ErrorTimeout    = ErrorCode{Code: 103, Msg: "request timeout"}
	ErrorRateLimit  = ErrorCode{Code: 104, Msg: "exceed quota limit"}
	ErrorInternal   = ErrorCode{Code: 105, Msg: "internal error"}

	// limit error code
	ErrorRateLimitByClient = ErrorCode{Code: 10000, Msg: "rate limited by client"}
	ErrorRateLimitByMethod = ErrorCode{Code: 10001, Msg: "rate limited by method"}
	ErrorRateLimitDefault  = ErrorCode{Code: 10002, Msg: "rate limited by default"}
)

func ErrorToErrorCode(err error) ErrorCode {
	ec, ok := err.(ErrorCode)
	if !ok {
		return ErrorInternal
	}
	return ec
}
func ErrorToCode(err error) int32 {
	ec, ok := err.(ErrorCode)
	if !ok {
		return int32(101)
	}
	return ec.GetErrCode()
}
