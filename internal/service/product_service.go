package service

import (
	"gin-rest-api/internal/model"
	"gin-rest-api/internal/repository"
)

type ProductService struct {
	repository *repository.ProductRepository
}

func NewProductService(repository *repository.ProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) GetAllProducts() []model.Product {
	return s.repository.GetAll()
}

func (s *ProductService) GetProductById(id int) (*model.Product, bool) {
	return s.repository.GetByID(id)
}

func (s *ProductService) CreateProduct(product model.Product) model.Product {
	return s.repository.Create(product)
}

func (s *ProductService) DeleteProduct(id int) bool {
	return s.repository.Delete(id)
}
