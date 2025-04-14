package models

type Task struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed" gorm:"default:false"`
	UserEmail   string `json:"user_email" gorm:"index"`
}
