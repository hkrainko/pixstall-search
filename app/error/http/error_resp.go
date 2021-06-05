package http

import (
	"net/http"
	error2 "pixstall-search/domain/error"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(err error) (int, interface{}) {
	if domainError, isError := err.(error2.DomainError); isError {
		switch domainError {
		case error2.UnAuthError:
			return http.StatusUnauthorized, ErrorResponse{
				Message: domainError.Error(),
			}
		case error2.NotFoundError:
			return http.StatusNotFound, ErrorResponse{
				Message: domainError.Error(),
			}
		case error2.BadRequestError:
			return http.StatusBadRequest, ErrorResponse{
				Message: domainError.Error(),
			}
		default:
			return http.StatusInternalServerError, ErrorResponse{
				Message: err.Error(),
			}
		}
	} else {
		return http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		}
	}

}