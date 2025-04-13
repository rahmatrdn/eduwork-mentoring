# Product Management REST API

Aplikasi REST API untuk manajemen produk menggunakan Go dengan struktur repository dan service pattern.

## Teknologi yang Digunakan

- Go (Golang)
- MySQL (Database)
- GORM (ORM)
- Gorilla Mux (Router)
- Godotenv (Environment Variables)

## Struktur Proyek

```tree
product-management/
├── config/
│   └── database.go          # Konfigurasi database
├── handler/
│   └── product_handler.go   # HTTP request handlers
├── helper/
│   └── response.go          # Response formatter
├── middleware/
│   └── logging.go          # HTTP middleware
├── models/
│   └── product.go          # Data models
├── repository/
│   └── product_repository.go # Database operations
├── service/
│   └── product_service.go   # Business logic
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── main.go
└── README.md
```

## Cara Menjalankan

1. Clone repository

```bash
git clone <https://github.com/rahmatrdn/eduwork-mentoring/tree/fani>
cd product-management
```

2. Copy file .env.example ke .env

```bash
cp .env.example .env
```

3. Sesuaikan konfigurasi database di file .env

4. Buat database MySQL

```sql
CREATE DATABASE product_db;
```

5. Install dependencies

```bash
go mod tidy
```

6. Jalankan aplikasi

```bash
go run main.go
```

## API Endpoints

### 1. Get All Products

- Method: GET
- URL: `/products`
- Response:

```json
{
  "status": true,
  "message": "Berhasil mengambil data produk",
  "data": [
    {
      "id": 1,
      "name": "Laptop Gaming",
      "description": "Laptop Gaming Terbaru",
      "price": 15000000,
      "stock": 10
    }
  ]
}
```

### 2. Get Product by ID

- Method: GET
- URL: `/products/{id}`
- Response:

```json
{
  "status": true,
  "message": "Berhasil mengambil data produk",
  "data": {
    "id": 1,
    "name": "Laptop Gaming",
    "description": "Laptop Gaming Terbaru",
    "price": 15000000,
    "stock": 10
  }
}
```

### 3. Create Product

- Method: POST
- URL: `/products`
- Body:

```json
{
  "name": "Laptop Gaming",
  "description": "Laptop Gaming Terbaru",
  "price": 15000000,
  "stock": 10
}
```

- Response:

```json
{
  "status": true,
  "message": "Berhasil membuat produk",
  "data": {
    "id": 1,
    "name": "Laptop Gaming",
    "description": "Laptop Gaming Terbaru",
    "price": 15000000,
    "stock": 10
  }
}
```

### 4. Update Product

- Method: PUT
- URL: `/products/{id}`
- Body:

```json
{
  "name": "Laptop Gaming Updated",
  "description": "Laptop Gaming Terbaru Update",
  "price": 16000000,
  "stock": 5
}
```

- Response:

```json
{
  "status": true,
  "message": "Berhasil mengupdate produk",
  "data": {
    "id": 1,
    "name": "Laptop Gaming Updated",
    "description": "Laptop Gaming Terbaru Update",
    "price": 16000000,
    "stock": 5
  }
}
```

### 5. Delete Product

- Method: DELETE
- URL: `/products/{id}`
- Response:

```json
{
  "status": true,
  "message": "Berhasil menghapus produk"
}
```

## Error Response Format

```json
{
  "status": false,
  "message": "Pesan error",
  "errors": ["Detail error"]
}
```

## Author

[Fani Andrianto]

## License

MIT
