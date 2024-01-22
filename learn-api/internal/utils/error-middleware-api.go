package utils

import (
	"learn-api/internal/model"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func validationError(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		response := model.APIResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		WriteToResponseBody(writer, response)
		return true
	}

	return false
}

func notFoundError(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := model.APIResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		WriteToResponseBody(writer, response)
		return true
	}

	return false
}

func internalServerError(writer http.ResponseWriter, _ *http.Request, err any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	response := model.APIResponse{
		Code:   http.StatusInternalServerError,
		Status: "Failure",
		Data:   err,
	}

	WriteToResponseBody(writer, response)
}

func ErrorHandlerHTTP(writer http.ResponseWriter, request *http.Request, err any) {
	if validationError(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				ErrorHandlerHTTP(writer, request, err)
				return
			}
		}()

		next.ServeHTTP(writer, request)
	})
}
