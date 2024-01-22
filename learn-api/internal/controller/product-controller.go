package controller

import (
	"learn-api/internal/model"
	"learn-api/internal/usecase"
	"learn-api/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductControllerImpl struct {
	usecase usecase.GenericUsecase[model.Product]
}

func New(usecase usecase.GenericUsecase[model.Product]) GenericController {
	return &ProductControllerImpl{
		usecase: usecase,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	payload := model.Product{}
	utils.ReadFromRequestBody(request, &payload)

	utils.WriteToResponseBody(writer, model.APIResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   controller.usecase.Create(payload),
	})
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	utils.WriteToResponseBody(writer, model.APIResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   controller.usecase.FindAll(),
	})
}

func (controller *ProductControllerImpl) FindByID(writer http.ResponseWriter, request *http.Request) {
	productIDLong, err := strconv.ParseUint(mux.Vars(request)["id"], 10, 32)
	utils.PanicIfError(err)

	utils.WriteToResponseBody(writer, model.APIResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   controller.usecase.FindByID(uint(productIDLong)),
	})
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	productIDLong, err := strconv.ParseUint(mux.Vars(request)["id"], 10, 32)
	utils.PanicIfError(err)

	controller.usecase.Delete(uint(productIDLong))
	utils.WriteToResponseBody(writer, model.APIResponse{
		Code:   http.StatusOK,
		Status: "Success",
	})
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	productIDLong, err := strconv.ParseUint(mux.Vars(request)["id"], 10, 32)
	utils.PanicIfError(err)

	payload := model.Product{}
	utils.ReadFromRequestBody(request, &payload)

	utils.WriteToResponseBody(writer, model.APIResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   controller.usecase.Update(uint(productIDLong), payload),
	})
}
