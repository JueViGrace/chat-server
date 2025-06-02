package types

import (
	"net/http"
	"time"
)

type APIResponse struct {
	Status      int       `json:"status"`
	Description string    `json:"description"`
	Data        any       `json:"data"`
	Message     string    `json:"message"`
	Time        time.Time `json:"time"`
}

func NewAPIResponse(status int, description string, data any, message string) *APIResponse {
	return &APIResponse{
		Status:      status,
		Description: description,
		Data:        data,
		Message:     message,
		Time:        time.Now(),
	}
}

func RespondOk(data any, message string) *APIResponse {
	return NewAPIResponse(http.StatusOK, http.StatusText(http.StatusOK), data, message)
}

func RespondCreated(data any, message string) *APIResponse {
	return NewAPIResponse(http.StatusCreated, http.StatusText(http.StatusCreated), data, message)
}

func RespondAccepted(data any, message string) *APIResponse {
	return NewAPIResponse(http.StatusAccepted, http.StatusText(http.StatusAccepted), data, message)
}

func RespondNoContent(data any, message string) *APIResponse {
	return NewAPIResponse(http.StatusNoContent, http.StatusText(http.StatusNoContent), data, message)
}

func RespondBadRequest(data any, message string) *APIResponse {
	return NewAPIResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), data, message)
}

func RespondUnauthorized(data any, message string) *APIResponse {
	return NewAPIResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), data, message)
}

func RespondForbbiden(data any, message string) *APIResponse {
	return NewAPIResponse(http.StatusForbidden, http.StatusText(http.StatusForbidden), data, message)
}

func RespondNotFound(data any, message string) *APIResponse {
	return NewAPIResponse(http.StatusNotFound, http.StatusText(http.StatusNotFound), data, message)
}

func RespondInternalServerError(data any, message string) *APIResponse {
	return NewAPIResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), data, message)
}
