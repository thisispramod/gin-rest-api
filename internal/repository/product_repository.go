package repository

import "gin-rest-api/internal/model"

type ProductRepository struct {
	products []model.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: []model.Product{
			{
				ID:    1,
				Name:  "Keyboard",
				Price: 1500,
			},
			{
				ID:    2,
				Name:  "Mouse",
				Price: 300,
			},
			{
				ID:    3,
				Name:  "Moniter",
				Price: 8300,
			},
		},
	}
}

func (r *ProductRepository) GetAll() []model.Product {
	return r.products
}

// Get Product by Id

func (r *ProductRepository) GetByID(id int) (*model.Product, bool) {
	for _, product := range r.products {
		if product.ID == id {
			return &product, true
		}
	}
	return nil, false
}

func (r *ProductRepository) Create(product model.Product) model.Product {
	r.products = append(r.products, product)
	return product
}

func (r *ProductRepository) Delete(id int) bool {
	for i, product := range r.products {
		if product.ID == id {
			r.products = append(
				r.products[:i],
				r.products[i+1:]...,
			)
			return true
		}
	}
	return false
}
