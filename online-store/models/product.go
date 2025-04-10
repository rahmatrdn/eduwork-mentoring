package models

import "time"

type Product struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Nama        string    `json:"nama"`
    Deskripsi   string    `json:"deskripsi"`
    Harga       float64   `json:"harga"`
    Stok        int       `json:"stok"`
    CreatedAt   time.Time `json:"created_at"`
}
