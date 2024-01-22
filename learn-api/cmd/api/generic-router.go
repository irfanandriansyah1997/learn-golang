package main

import (
	"fmt"
	"learn-api/internal/controller"
	"learn-api/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GenericRouter[T any] struct {
	baseURLPath string
	controller  *controller.GenericController
}

func (router *GenericRouter[T]) handle(writer http.ResponseWriter, request *http.Request) {
	idLong, err := strconv.ParseUint(mux.Vars(request)["id"], 10, 32)

	if request.URL.EscapedPath() != router.baseURLPath && err != nil {
		utils.PanicIfError(err)
		return
	}

	switch request.Method {
	case http.MethodGet:
		if idLong != 0 {
			(*(router.controller)).FindByID(writer, request)
		} else {
			(*(router.controller)).FindAll(writer, request)
		}

	case http.MethodPost:
		(*(router.controller)).Create(writer, request)

	case http.MethodPut:
		(*(router.controller)).Update(writer, request)

	case http.MethodDelete:
		(*(router.controller)).Delete(writer, request)

	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (router *GenericRouter[T]) registerRoutes(mux *mux.Router) {
	mux.HandleFunc(router.baseURLPath, router.handle)
	mux.HandleFunc(fmt.Sprintf("%v/{id}", router.baseURLPath), router.handle)
}

func NewGenericRouter[T any](baseUrlPath string, mux *mux.Router, controller *controller.GenericController) *GenericRouter[T] {
	router := GenericRouter[T]{
		baseURLPath: baseUrlPath,
		controller:  controller,
	}

	router.registerRoutes(mux)
	return &router
}
