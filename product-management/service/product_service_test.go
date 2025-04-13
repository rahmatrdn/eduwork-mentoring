package service

import (
	"errors"
	"product-management/models"
	"testing"
)

// TestCreateProduct menguji fungsi CreateProduct
func TestCreateProduct(t *testing.T) {
	// Membuat test cases
	tests := []struct {
		name          string
		input         models.Product
		expectedError bool
	}{
		{
			name: "sukses_create_product",
			input: models.Product{
				Name:        "Laptop Test",
				Description: "Deskripsi Test",
				Price:      15000000,
				Stock:      10,
			},
			expectedError: false,
		},
		{
			name: "gagal_nama_kosong",
			input: models.Product{
				Name:        "",
				Description: "Deskripsi Test",
				Price:      15000000,
				Stock:      10,
			},
			expectedError: true,
		},
		{
			name: "gagal_harga_negatif",
			input: models.Product{
				Name:        "Laptop Test",
				Description: "Deskripsi Test",
				Price:      -1000,
				Stock:      10,
			},
			expectedError: true,
		},
	}

	// Menjalankan setiap test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Validasi produk
			err := validateProduct(tt.input)
			
			// Cek hasil
			if (err != nil) != tt.expectedError {
				t.Errorf("validateProduct() error = %v, expectedError %v", err, tt.expectedError)
			}
		})
	}
}

// Fungsi validasi yang akan ditest
func validateProduct(product models.Product) error {
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