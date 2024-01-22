package main

import (
	"fmt"
	productController "learn-api/internal/controller"
	"learn-api/internal/model"
	productRepo "learn-api/internal/repository"
	productUsecase "learn-api/internal/usecase"
	"learn-api/internal/utils"

	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func main() {
	utils.LoadAppConfig("env", "app", ".")

	validate := validator.New()
	router := mux.NewRouter().StrictSlash(true)

	RegisterProductRoutes(router, validate)
	router.Use(utils.ErrorMiddleware)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", utils.AppConfig.Port))
	server := http.Server{
		Addr:              fmt.Sprintf("localhost:%v", utils.AppConfig.Port),
		Handler:           router,
		ReadHeaderTimeout: 2 * time.Second,
	}

	err := server.ListenAndServe()
	utils.PanicIfError(err)
}

func RegisterProductRoutes(router *mux.Router, validator *validator.Validate) {
	var productRepo = productRepo.New()
	var productUsecase = productUsecase.New(productRepo, validator)
	var productController = productController.New(productUsecase)
	NewGenericRouter[model.Product]("/api/products", router, &productController)
}
