package repository

import (
	"fmt"

	"learn-api/internal/model"
	"learn-api/internal/utils"

	"reflect"
)

type ProductRepo struct {
	products []model.Product
}

func New() GenericRepo[model.Product] {
	return &ProductRepo{products: make([]model.Product, 0)}
}

func (p *ProductRepo) Create(product model.Product) model.Product {
	newItem := product
	newItem.ID = uint(len(p.products)) + 1
	p.products = append(p.products, newItem)

	return newItem
}

func (p *ProductRepo) FindAll() []model.Product {
	return p.products
}

func (p *ProductRepo) FindByID(id uint) model.Product {
	product := model.Product{}

	for _, item := range p.products {
		if item.ID == id {
			product = item
		}
	}

	if reflect.ValueOf(product).IsZero() {
		utils.PanicIfNotFoundError(fmt.Errorf("Key '%d' not found", id))
	}

	return product
}

func (p *ProductRepo) Update(id uint, payload model.Product) model.Product {
	product := model.Product{}
	for index, item := range p.products {
		if item.ID == id {
			payload.ID = id

			p.products = append(p.products[:index], p.products[index+1:]...)
			p.products = append(p.products, payload)
			product = payload
		}
	}

	if reflect.ValueOf(product).IsZero() {
		utils.PanicIfNotFoundError(fmt.Errorf("Key '%d' not found", id))
	}

	return product
}

func (p *ProductRepo) Delete(id uint) {
	isAvailable := false
	for index, item := range p.products {
		if item.ID == id {
			p.products = append(p.products[:index], p.products[index+1:]...)
			isAvailable = true
		}
	}

	if !isAvailable {
		utils.PanicIfNotFoundError(fmt.Errorf("Key '%d' not found", id))
	}
}
