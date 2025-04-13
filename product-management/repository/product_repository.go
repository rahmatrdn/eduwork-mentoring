package repository

import (
	"product-management/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id uint) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *productRepository) Create(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *productRepository) Update(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
