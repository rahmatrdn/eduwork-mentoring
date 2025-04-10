package controllers

import (
	"net/http"
	"online-store/config"
	"online-store/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
    var input models.Product

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product := models.Product{
        Nama:      input.Nama,
        Deskripsi: input.Deskripsi,
        Harga:     input.Harga,
        Stok:      input.Stok,
        CreatedAt: input.CreatedAt,
    }

    if err := config.DB.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat produk"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Produk berhasil ditambahkan",
        "product": product,
    })
}

func GetAllProducts(c *gin.Context) {
    var products []models.Product

    if err := config.DB.Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data produk"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "products": products,
    })
}

func GetProductByID(c *gin.Context) {
    id := c.Param("id")
    var product models.Product

    if err := config.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"product": product})
}

func UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product

    // Cari dulu produknya
    if err := config.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
        return
    }

    // Bind input baru
    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update data
    product.Nama = input.Nama
    product.Deskripsi = input.Deskripsi
    product.Harga = input.Harga
    product.Stok = input.Stok

    if err := config.DB.Save(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate produk"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Produk berhasil diupdate",
        "product": product,
    })
}

func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product

    // Cek dulu apakah produk ada
    if err := config.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
        return
    }

    // Hapus produk
    if err := config.DB.Delete(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus produk"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Produk berhasil dihapus",
    })
}
