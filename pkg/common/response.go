package common

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DebugType struct {
	Property   string
	Error      error
	Additional string
	Function   string
}

func ErrorResponse(
	code int,
	message string,
	err error,
) Response {
	return Response{
		Data:    json.NewEncoder(nil),
		Code:    code,
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

func SuccessResponse(message string, code int, data interface{}) Response {
	return Response{
		Data:    data,
		Code:    code,
		Message: message,
		Status:  http.StatusOK,
	}
}
