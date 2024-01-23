package entities

import "net/http"

type GenericHTTPHandler interface {
	Create(writer http.ResponseWriter, r *http.Request)
	FindAll(writer http.ResponseWriter, r *http.Request)
	FindById(writer http.ResponseWriter, r *http.Request)
	Delete(writer http.ResponseWriter, r *http.Request)
	Update(writer http.ResponseWriter, r *http.Request)
}
