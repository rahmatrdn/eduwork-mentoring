package service

import (
	"errors"
	"product-management/models"
	"product-management/repository"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(id uint, product models.Product) (models.Product, error)
	DeleteProduct(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetProductByID(id uint) (models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) CreateProduct(product models.Product) (models.Product, error) {
	if product.Name == "" {
		return models.Product{}, errors.New("nama produk tidak boleh kosong")
	}
	if product.Price <= 0 {
		return models.Product{}, errors.New("harga produk harus lebih dari 0")
	}
	return s.repo.Create(product)
}

func (s *productService) UpdateProduct(id uint, product models.Product) (models.Product, error) {
	existingProduct, err := s.repo.FindByID(id)
	if err != nil {
		return models.Product{}, errors.New("produk tidak ditemukan")
	}
	
	product.ID = existingProduct.ID
	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("produk tidak ditemukan")
	}
	return s.repo.Delete(id)
}

func (s *productService) validateProduct(product models.Product) error {
    if product.Name == "" {
        return errors.New("nama produk tidak boleh kosong")
    }
    if product.Price <= 0 {
        return errors.New("harga produk harus lebih dari 0")
    }
    if product.Stock < 0 {
        return errors.New("stok tidak boleh negatif")
    }
    return nil
}