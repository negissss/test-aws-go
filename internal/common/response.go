package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Error      bool        `json:"error"`
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func Respond(context *gin.Context, error bool, statusCode int, message string, data interface{}) {

	isEmpty := data == nil

	if !error {
		if !isEmpty {
			context.JSON(statusCode, APIResponse{
				Error:      false,
				StatusCode: statusCode,
				Message:    message,
				Data:       data,
			})
		} else {
			context.JSON(statusCode, APIResponse{
				Error:      false,
				StatusCode: statusCode,
				Message:    message,
			})
		}
		return
	}
	context.JSON(statusCode, APIResponse{
		Error:      true,
		StatusCode: statusCode,
		Message:    message,
	})

}

func SuccessResponse(context *gin.Context, message string, data interface{}) {
	Respond(context, false, http.StatusOK, message, data)
}

func ErrorResponse(context *gin.Context, statusCode int, message string) {
	Respond(context, true, statusCode, message, nil)
}

type DateCount struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

type ChainTxCount struct {
	TargetChain string `json:"target_chain"`
	Count       int64  `json:"count"`
}
