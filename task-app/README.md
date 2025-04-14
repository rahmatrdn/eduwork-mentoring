# Task Management API

Sistem ini adalah RESTful API sederhana untuk mengelola tugas (Task) menggunakan Golang, Gin, dan GORM.

## 📦 Fitur

- Registrasi User
- Login User (dengan JWT)
- CRUD Task (Create, Read, Update, Delete)
- Tandai Task sebagai Selesai

## 🧱 Tech Stack

- Golang
- Gin
- GORM
- MySQL
- JWT Authentication
- Postman (untuk dokumentasi dan testing)

## ⚙️ Instalasi & Setup

### 1. Clone Repo & Masuk Folder

```bash
git clone https://github.com/username/task-app.git
cd task-app
```

### 2. Buat File `.env`

```
DB_USER=root
DB_PASS=
DB_NAME=taskdb
DB_HOST=127.0.0.1
DB_PORT=3306
JWT_SECRET=rahasia123
```

### 3. Jalankan MySQL & Buat DB

```sql
CREATE DATABASE taskdb;
```

### 4. Jalankan Aplikasi

```bash
go mod tidy
go run main.go
```

Server akan jalan di `http://localhost:8080`

## 🔑 Autentikasi

Gunakan endpoint `/api/login` untuk mendapatkan JWT token.
Token ini digunakan untuk akses endpoint yang dilindungi seperti `/api/tasks`.

Tambahkan ke header:

```
Authorization: Bearer <token>
```

## 📮 Endpoint API

| Method | Endpoint                  | Deskripsi              |
| ------ | ------------------------- | ---------------------- |
| POST   | `/api/register`           | Registrasi user        |
| POST   | `/api/login`              | Login dan dapatkan JWT |
| GET    | `/api/tasks`              | List semua task        |
| GET    | `/api/tasks/:id`          | Detail task tertentu   |
| POST   | `/api/tasks`              | Tambah task baru       |
| PUT    | `/api/tasks/:id`          | Update task            |
| DELETE | `/api/tasks/:id`          | Hapus task             |
| PUT    | `/api/tasks/:id/complete` | Tandai task selesai    |

## 📬 Testing via Postman

1. Import Collection & Environment (file JSON sudah tersedia).
2. Jalankan `register` lalu `login` untuk mendapatkan token.
3. Gunakan token untuk akses endpoint `tasks`.

## 🙌 Author

- Nama: Fani
- Eduwork Golang Mentoring
