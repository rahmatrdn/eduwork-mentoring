Sistem Manajemen Pengguna (User Management System)

# 🧑‍💼 API Sistem Manajemen Pengguna

Sebuah RESTful API sederhana dan aman untuk manajemen pengguna, dibangun menggunakan **Golang**, **Gin**, **GORM**, **MySQL**, dan **JWT**.

## 🚀 Fitur

- ✅ Pendaftaran Pengguna (Sign Up)
- ✅ Login dengan Autentikasi JWT
- ✅ Melihat Profil Pengguna
- ✅ Update Profil Pengguna
- ✅ Ganti Kata Sandi (dengan autentikasi)
- ✅ Lupa Kata Sandi (mengirim token ke email)
- ✅ Reset Kata Sandi dengan Token
- ✅ Enkripsi kata sandi dengan bcrypt
- ✅ Middleware untuk otorisasi akses
- ✅ Penyimpanan data menggunakan MySQL

## 🛠 Teknologi yang Digunakan

- Go (Golang)
- Gin Gonic (Framework HTTP)
- GORM (ORM)
- MySQL
- JWT (Autentikasi & Otorisasi)
- Postman (untuk testing API)

## 📁 Struktur Folder

user-management/ │ ├── config/ # Konfigurasi koneksi database ├── controllers/ # Logika utama aplikasi (handler) ├── models/ # Struktur data model pengguna ├── routes/ # Rute-rute API ├── middleware/ # Middleware JWT ├── utils/ # Fungsi pembantu (helper) & token ├── main.go # Entry point aplikasi

## ⚙️ Cara Menjalankan Aplikasi

### 1. Install dependency Go:

```bash
go mod tidy
2. Buat database MySQL:
CREATE DATABASE user_management;
3. Atur konfigurasi koneksi di config/db.go:
dsn := "root:password_mysql@tcp(127.0.0.1:3306)/user_management?parseTime=true"
4. Jalankan aplikasi:
go run main.go

🔌 Daftar Endpoint API

Method	Endpoint	Auth	Deskripsi
POST	/register	❌	Daftar pengguna baru
POST	/login	❌	Login dan dapatkan token JWT
GET	/user	✅	Lihat profil pengguna saat ini
PUT	/user	✅	Update data profil pengguna
PUT	/change-password	✅	Ganti kata sandi (dalam keadaan login)
POST	/forgot-password	❌	Minta token reset kata sandi
POST	/reset-password	❌	Reset kata sandi menggunakan token

🔐 Autentikasi JWT
Gunakan token JWT di header setiap request yang membutuhkan autentikasi:

makefile
Salin
Edit
Authorization: Bearer <eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhbmlAZXhhbXBsZS5jb20iLCJleHAiOjE3NDQ5NjI0ODEsInVzZXJfaWQiOjF9.EWHcQ0Mxvkld_3RwE7r5TiCnj5oifTYz1VQpak6S8iA>
👨‍💻 Tentang Proyek Ini
Dibuat dengan ❤️ menggunakan Golang sebagai bagian dari tugas mentoring.
Nama: [Fani Andrianto]
Platform: [Eduwork | Bootcamp]
```
