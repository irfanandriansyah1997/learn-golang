package usecase

import (
	"learn-api/internal/model"
	"learn-api/internal/repository"
	"learn-api/internal/utils"

	"github.com/go-playground/validator/v10"
)

type ProductUsecaseImpl struct {
	repo     repository.GenericRepo[model.Product]
	validate *validator.Validate
}

func New(repo repository.GenericRepo[model.Product], validate *validator.Validate) GenericUsecase[model.Product] {
	return &ProductUsecaseImpl{
		repo:     repo,
		validate: validate,
	}
}

func (usecase *ProductUsecaseImpl) Create(payload model.Product) model.Product {
	err := usecase.validate.Struct(payload)
	utils.PanicIfError(err)

	return usecase.repo.Create(payload)
}

func (usecase *ProductUsecaseImpl) FindAll() []model.Product {
	return usecase.repo.FindAll()
}

func (usecase *ProductUsecaseImpl) FindByID(id uint) model.Product {
	return usecase.repo.FindByID(id)
}

func (usecase *ProductUsecaseImpl) Delete(id uint) {
	// INFO: Find product by ID, if the product not available automaticaly invoked panic handler
	usecase.repo.FindByID(id)

	usecase.repo.Delete(id)
}

func (usecase *ProductUsecaseImpl) Update(id uint, payload model.Product) model.Product {
	err := usecase.validate.Struct(payload)
	utils.PanicIfError(err)

	// INFO: Find product by ID, if the product not available automaticaly invoked panic handler
	usecase.repo.FindByID(id)

	return usecase.repo.Update(id, payload)
}
