package controllers

import (
	"net/http"
	"online-store/config"
	"online-store/models"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
    var input struct {
        Username string `json:"username" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Cek apakah email sudah terdaftar
    var existingUser models.User
    if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah terdaftar"})
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
        return
    }

    // Buat user
    user := models.User{
        Username:  input.Username,
        Email:     input.Email,
        Password:  string(hashedPassword),
        CreatedAt: time.Now(),
    }

    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat user"})
        return
    }

    // Response tanpa password
	c.JSON(http.StatusCreated, gin.H{
        "message": "Registrasi berhasil",
        "user": gin.H{
            "id":         user.ID,
            "username":   user.Username,
            "email":      user.Email,
            "created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
        },
    })
}

func LoginUser(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
        return
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Login berhasil",
        "user": gin.H{
            "id":         user.ID,
            "username":   user.Username,
            "email":      user.Email,
            "created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
        },
    })
}
