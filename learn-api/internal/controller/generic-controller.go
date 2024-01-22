package controller

import "net/http"

type GenericController interface {
	Create(writer http.ResponseWriter, r *http.Request)
	FindAll(writer http.ResponseWriter, r *http.Request)
	FindByID(writer http.ResponseWriter, r *http.Request)
	Delete(writer http.ResponseWriter, r *http.Request)
	Update(writer http.ResponseWriter, r *http.Request)
}
