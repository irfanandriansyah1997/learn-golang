package repositories

import (
	"context"
	"fmt"
	"learn-grpc/internal/entities"
	"learn-grpc/internal/utils"

	"github.com/google/uuid"
)

func generatePhotos(photos []string) []entities.Asset {
	formattedPhotos := make([]entities.Asset, 0)
	for _, photo := range photos {
		asset, err := utils.GenerateOptimizeAsset(photo)

		if err == nil {
			formattedPhotos = append(formattedPhotos, *asset)
		}
	}

	return formattedPhotos
}

type _ProductRepoImpl struct {
	products []entities.Product
}

func NewProductRepo() entities.GenericRepo[entities.ProductRequest, entities.Product, string] {
	return &_ProductRepoImpl{
		products: make([]entities.Product, 0),
	}
}

func (p *_ProductRepoImpl) Create(_ context.Context, req entities.ProductRequest) entities.Product {
	newProduct := entities.Product{}
	newProduct.ID = uuid.New().String()
	newProduct.Description = req.Description
	newProduct.Name = req.Name
	newProduct.Photos = generatePhotos(req.Photos)

	if formattedPrice, err := utils.FormattingPrice(req.Price); err != nil && formattedPrice != nil {
		newProduct.Price = *formattedPrice
	}

	p.products = append(p.products, newProduct)

	return newProduct
}

func (p *_ProductRepoImpl) FindAll(_ context.Context) []entities.Product {
	return p.products
}

func (p *_ProductRepoImpl) FindByID(_ context.Context, id string) (*entities.Product, error) {
	var product *entities.Product

	for _, item := range p.products {
		if item.ID == id {
			product = &item
		}
	}

	if product != nil {
		return product, nil
	}

	return nil, utils.NewNotFoundError(fmt.Sprintf("product %s is not found", id))
}

func (p *_ProductRepoImpl) Delete(_ context.Context, id string) bool {
	isAvailable := false

	for index, item := range p.products {
		if item.ID == id {
			p.products = append(p.products[:index], p.products[index+1:]...)
			isAvailable = true
		}
	}

	if !isAvailable {
		utils.PanicIfNotFoundError(fmt.Errorf("product %s is not found", id))
		return false
	}

	return true
}

func (p *_ProductRepoImpl) Update(_ context.Context, id string, req entities.ProductRequest) *entities.Product {
	var selectedProduct *entities.Product

	for index, item := range p.products {
		if item.ID == id {
			selectedProduct = &item
			selectedProduct.ID = id
			selectedProduct.Description = req.Description
			selectedProduct.Name = req.Name
			selectedProduct.Photos = generatePhotos(req.Photos)

			if formattedPrice, err := utils.FormattingPrice(req.Price); err != nil && formattedPrice != nil {
				selectedProduct.Price = *formattedPrice
			}

			p.products = append(p.products[:index], p.products[index+1:]...)
			p.products = append(p.products, *selectedProduct)
		}
	}

	if selectedProduct == nil {
		utils.PanicIfNotFoundError(fmt.Errorf("product %s is not found", id))
	}

	return selectedProduct
}
