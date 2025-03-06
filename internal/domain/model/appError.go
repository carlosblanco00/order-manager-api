package model

import (
	"errors"
	"fmt"
	"log"
)

const (
	ErrCodeInternalServerError = "internal_server_error"
	ErrCodeInvalidParams       = "invalid_params"
	ErrCodeNotFound            = "not_found"
	ErrCodeBadRequest          = "bad_request"
	ErrCodeIdenpotency         = "Idempotency_Key_already_exist"
	ErrCodeHeaderIdenpotency   = "Header_Idempotency_Key_missing"
	ErrCodeStock               = "insufficient_stock"
)

var (
	ErrIdenpotency       = errors.New("the Idempotency_Key already exist")
	ErrIncorrectID       = errors.New("incorrect id error")
	ErrNotFound          = errors.New("record not found error")
	ErrBadRequest        = errors.New("bad request")
	ErrStock             = errors.New("insufficient stock for product")
	ErrHeaderIdenpotency = errors.New("Idempotency-Key header is missing")
)

type AppError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func NewAppError(code string, msg string) AppError {
	return AppError{
		Code: code,
		Msg:  msg,
	}
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Msg)
}

func ManageError(err error) error {
	var appErr AppError

	switch {
	case errors.Is(err, ErrIdenpotency):
		log.Println("error Idempotency-Key")
		appErr = AppError{
			Code: ErrCodeIdenpotency,
			Msg:  "Error creating order: Idempotency-Key already exist",
		}
	case errors.Is(err, ErrIncorrectID):
		log.Println("incorrect id error")
		appErr = AppError{
			Code: ErrCodeInvalidParams,
			Msg:  "Incorrect id",
		}
	case errors.Is(err, ErrNotFound):
		log.Println("not found error")
		appErr = AppError{
			Code: ErrCodeNotFound,
			Msg:  "Not found",
		}
	case errors.Is(err, ErrBadRequest):
		log.Println("bad request error")
		appErr = AppError{
			Code: ErrCodeBadRequest,
			Msg:  "Invalid request",
		}
	case errors.Is(err, ErrHeaderIdenpotency):
		log.Println("Idempotency-Key header is missing")
		appErr = AppError{
			Code: ErrCodeHeaderIdenpotency,
			Msg:  "Idempotency-Key header is missing",
		}
	case errors.Is(err, ErrStock):
		log.Println("insufficient stock for product")
		appErr = AppError{
			Code: ErrCodeStock,
			Msg:  "insufficient stock for product",
		}
	default:
		log.Println(err.Error())
		appErr = AppError{
			Code: ErrCodeInternalServerError,
			Msg:  "Server Error",
		}
	}

	return appErr
}
