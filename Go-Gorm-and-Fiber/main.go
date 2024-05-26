package main

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       uint      `json:"price"`
	Stock       uint      `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	// Connect to database
	dsn := "root:@tcp(localhost:3306)/toko_online?charset=utf8mb4&parseTime=true&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Product{})

	// Create a Fiber app
	app := fiber.New()

	// Routes
	app.Post("/api/users/register", registerUser)
	app.Post("/api/users/login", authenticateUser)

	app.Post("/api/products", createProduct)
	app.Get("/api/products", getProducts)
	app.Get("/api/products/:id", getProduct)
	app.Put("/api/products/:id", updateProduct)
	app.Delete("/api/products/:id", deleteProduct)

	// Start server
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}

// User handlers
func registerUser(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	user.Password = string(hashedPassword)

	db.Create(&user)
	return c.Status(http.StatusCreated).JSON(user)
}

func authenticateUser(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var existingUser User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Authentication successful"})
}

// Product handlers
func createProduct(c *fiber.Ctx) error {
	var product Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	db.Create(&product)
	return c.Status(http.StatusCreated).JSON(product)
}

func getProducts(c *fiber.Ctx) error {
	var products []Product
	db.Find(&products)
	return c.Status(http.StatusOK).JSON(products)
}

func getProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.Status(http.StatusOK).JSON(product)
}

func updateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	db.Save(&product)
	return c.Status(http.StatusOK).JSON(product)
}

func deleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	db.Delete(&product)
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Product deleted successfully"})
}
