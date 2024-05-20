package models

import (
	"github.com/labstack/echo/v4"
)

// Response Base response
type Response struct {
	StatusCode int            `json:"-"`
	Success    bool           `json:"success"`
	Message    string         `json:"message,omitempty"`
	Data       map[string]any `json:"data,omitempty"`
}

func (response *Response) SendResponse(c echo.Context) error {
	return c.JSON(response.StatusCode, response)
}

//func SendResponseData(c echo.Context, data map[string]any) error {
//	response := &Response{
//		StatusCode: http.StatusOK,
//		Success:    true,
//		Data:       data,
//	}
//	return response.SendResponse(c)
//}
//
//func SendErrorResponse(c echo.Context, status int, message string) error {
//	response := &Response{
//		StatusCode: status,
//		Success:    false,
//		Message:    message,
//	}
//	return response.SendResponse(c)
//}