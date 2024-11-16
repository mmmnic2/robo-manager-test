package dto

import "fmt"

type Response struct {
	StatusCode int         `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
type ErrorResponse struct {
	StatusCode int    `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (r ErrorResponse) Error() string {
	return fmt.Sprintf("Error %d: %s", r.StatusCode, r.Message)
}
