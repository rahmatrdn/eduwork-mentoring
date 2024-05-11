package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mahasiswa struct {
	Nama    string `json:"nama"`
	Nim     int    `json:"nim"`
	Jurusan string `json:"jurusan"`
}

func main() {
	dsn := "root:@tcp(localhost:3306)/mahasiswa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	app := fiber.New()

	// add mahasiswa
	app.Post("/dataMahasiswa", func(c *fiber.Ctx) error {
		var mahasiswa Mahasiswa
		if err := c.BodyParser(&mahasiswa); err != nil {
			return err
		}
		db.Create(&mahasiswa)
		return c.JSON(mahasiswa)
	})

	// delete mahasiswa
	app.Delete("/dataMahasiswa/:nim", func(c *fiber.Ctx) error {
		nim := c.Params("nim")

		// Convert nim to int
		nimInt, err := strconv.Atoi(nim)
		if err != nil {
			return err
		}

		if err := db.Where("nim = ?", nimInt).Delete(&Mahasiswa{}).Error; err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusNoContent)
	})

	// Show mahasiswa
	app.Get("/dataMahasiswa", func(c *fiber.Ctx) error {
		var dataMahasiswa []Mahasiswa
		if err := db.Find(&dataMahasiswa).Error; err != nil {
			return err
		}
		return c.JSON(dataMahasiswa)
	})

	port := "3000"
	app.Listen(":" + port)
}
