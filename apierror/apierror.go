package apierror

import "net/http"

type ApiError interface {
	Code() string
	Domain() string
	Message() string
	Status() int
}

type apiError struct {
	ErrorCode    string `json:"code"`
	ErrorDomain  string `json:"domain"`
	ErrorMessage string `json:"message"`
	ErrorStatus  int    `json:"status"`
}

func (e apiError) Message() string {
	return e.ErrorMessage
}

func (e apiError) Code() string {
	return e.ErrorCode
}

func (e apiError) Domain() string {
	return e.ErrorDomain
}

func (e apiError) Status() int {
	return e.ErrorStatus
}

func NewApiError(errorCode string, errorDomain string, errorMessage string, errorStatus int) ApiError {
	return apiError{errorCode, errorDomain, errorMessage, errorStatus}
}

func NewNotFoundApiError(errorDomain string, errorMessage string) ApiError {
	return apiError{"not_found", errorDomain, errorMessage, http.StatusNotFound}
}

func NewBadRequestApiError(errorDomain string, errorMessage string) ApiError {
	return apiError{"bad_request", errorDomain, errorMessage, http.StatusBadRequest}
}

func NewInternalServerApiError(errorDomain string, errorMessage string) ApiError {
	return apiError{"internal_server_error", errorDomain, errorMessage, http.StatusInternalServerError}
}

func NewUnauthorizedApiError(errorDomain string, errorMessage string) ApiError {
	return apiError{"unauthorized", errorDomain, errorMessage, http.StatusUnauthorized}
}
