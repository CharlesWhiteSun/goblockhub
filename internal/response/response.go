package response

import (
	"net/http"

	"github.com/CharlesWhiteSun/gomodx/errorx"
	"github.com/CharlesWhiteSun/gomodx/logger"

	"github.com/gin-gonic/gin"
)

type IResponseHandler interface {
	Success(c *gin.Context, data any, msg string)
	Error(c *gin.Context, code errorx.ErrorCode, msg ...string)
}

type ResponseHandler struct{}

func NewResponseHandler() IResponseHandler {
	return &ResponseHandler{}
}

func (r *ResponseHandler) Success(c *gin.Context, data any, msg string) {
	logger.Infof("API Success| Path: %s| Msg: %s| Data: %+v", c.FullPath(), msg, data)
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    errorx.SUCCESS,
		"message": msg,
		"data":    data,
	})
}

func (r *ResponseHandler) Error(c *gin.Context, code errorx.ErrorCode, msg ...string) {
	message := code.String()
	if len(msg) > 0 {
		message = msg[0]
	}

	logger.Errorf("API Error| Path: %s| Code: %d| Msg: %s", c.FullPath(), code, message)
	
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"code":    code,
		"message": message,
		"data":    nil,
	})
}
