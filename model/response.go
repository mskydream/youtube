package model

type ErrorResponse struct {
	IsSuccess bool
	Message   string
}

type SuccessResponse struct {
	IsSuccess bool
	Message   string
	Data      interface{}
}
