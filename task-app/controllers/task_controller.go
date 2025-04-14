package controllers

import (
	"net/http"

	"task-app/config"
	"task-app/models"

	"github.com/gin-gonic/gin"
)

// Buat task baru
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil email user dari token
	userEmail, exists := c.Get("userEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Set email user ke task
	task.UserEmail = userEmail.(string)

	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task berhasil dibuat", "task": task})
}

// Ambil semua task
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	userEmail, exists := c.Get("userEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		return
	}

	if err := config.DB.Where("user_email = ?", userEmail).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil task"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// Ambil task berdasarkan ID
func GetTaskByID(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	userEmail, exists := c.Get("userEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		return
	}

	if err := config.DB.Where("id = ? AND user_email = ?", id, userEmail).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Update task
func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	userEmail, exists := c.Get("userEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Cek apakah task ada dan milik user yang bersangkutan
	if err := config.DB.Where("id = ? AND user_email = ?", id, userEmail).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task tidak ditemukan"})
		return
	}

	// Update task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pastikan user_email tidak berubah
	task.UserEmail = userEmail.(string)

	if err := config.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task berhasil diupdate", "task": task})
}

// Hapus task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	userEmail, exists := c.Get("userEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		return
	}

	result := config.DB.Where("id = ? AND user_email = ?", id, userEmail).Delete(&models.Task{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus task"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task berhasil dihapus"})
}

// Tandai task sebagai selesai
func MarkTaskComplete(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	userEmail, exists := c.Get("userEmail")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Cek apakah task ada dan milik user yang bersangkutan
	if err := config.DB.Where("id = ? AND user_email = ?", id, userEmail).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task tidak ditemukan"})
		return
	}

	// Update status task
	task.IsCompleted = true
	if err := config.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate status task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task berhasil ditandai selesai", "task": task})
}
