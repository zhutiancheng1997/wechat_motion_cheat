package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const ginErrorCodeKey = "error_code"

// CommonResponse common response
type CommonResponse struct {
	*ErrorCode
	Data interface{} `json:"data"`
}

// SetSuccess set success
func (resp *CommonResponse) SetSuccess(ctx *gin.Context, data interface{}) {
	resp.setErrorCode(ctx, Ok, data)
}

// SetErrorCode set error code
func (resp *CommonResponse) SetErrorCode(ctx *gin.Context, errorCode ErrorCode) {
	resp.setErrorCode(ctx, errorCode, map[string]interface{}{})
}

func (resp *CommonResponse) setErrorCode(ctx *gin.Context, errorCode ErrorCode, data interface{}) {
	resp.ErrorCode = &errorCode
	resp.Data = data
	ctx.Set(ginErrorCodeKey, errorCode.GetErrCode())
}

func EchoErrorResponse(ctx *gin.Context, code ErrorCode) {
	resp := new(CommonResponse)
	resp.SetErrorCode(ctx, code)
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func EchoSuccessResponse(ctx *gin.Context, data interface{}) {
	resp := new(CommonResponse)
	resp.SetSuccess(ctx, data)
	ctx.JSON(http.StatusOK, resp)
}

func GetErrorCodeFromContext(ctx *gin.Context) int32 {
	code, exists := ctx.Get(ginErrorCodeKey)
	if !exists {
		return ErrorUnknown.GetErrCode()
	}
	return code.(int32)
}
