## 📦 Logistics System

Sistem backend manajemen logistik dengan layanan terpisah 

* `auth-service`: Autentikasi pengguna menggunakan JWT
* `logistics-service`: Manajemen pengiriman & tarif kurir

Dibangun dengan Golang (Gin, GORM) dan PostgreSQL.

---

## 🧱 Project Structure

```
logistics-system/
├── auth-service/         # Service untuk autentikasi user
├── logistics-service/    # Service untuk fitur logistik (shipment, kurir)
├── docs/                 # Swagger API specs
├── docker-compose.yml    # Multi-service orchestration
└── README.md
```

---

## 🚀 Getting Started

### Prasyarat:

* Docker & Docker Compose
* Go (jika ingin menjalankan lokal per service)

### Jalankan Semua Service:

```bash
docker-compose up --build
```

> Ini akan menjalankan:
>
> * PostgreSQL
> * Auth Service di port `8001`
> * Logistics Service di port `8002`

---

## 📌 Endpoints Overview

### 🛡️ Auth Service (`localhost:8001`)

| Method | Endpoint        | Deskripsi                          |
| ------ | --------------- | ---------------------------------- |
| POST   | `/register`     | Daftar user baru                   |
| POST   | `/login`        | Login & dapatkan JWT               |
| GET    | `/me`           | Profil user dari JWT               |
| GET    | `/verify-token` | Verifikasi JWT dan tampilkan claim |

### 📦 Logistics Service (`localhost:8002`)

| Method | Endpoint                | Deskripsi                                    |
| ------ | ----------------------- | -------------------------------------------- |
| POST   | `/shipments`            | Buat pengiriman barang                       |
| PUT    | `/shipments/:id/status` | Update status pengiriman                     |
| GET    | `/shipments/track`      | Lacak pengiriman berdasarkan nomor resi      |
| GET    | `/shipments`            | Ambil semua pengiriman oleh user             |
| GET    | `/courier-rates`        | Ambil seluruh tarif kurir                    |
| GET    | `/courier-rates/route`  | Ambil tarif berdasarkan origin & destination |

---

## 🧪 Testing

> Coming soon — unit test dengan `testify`.

---

## 📚 API Documentation

> Dokumentasi Swagger tersedia setelah service berjalan:

* [`localhost:8001/swagger/index.html`](http://localhost:8001/swagger/index.html) — Auth Service
* [`localhost:8002/swagger/index.html`](http://localhost:8002/swagger/index.html) — Logistics Service

---

## ⚙️ Teknologi yang Digunakan

* Language: **Go**
* Framework: [Gin](https://github.com/gin-gonic/gin)
* ORM: [GORM](https://gorm.io)
* Auth: JWT (via `golang-jwt/jwt/v5`)
* DB: PostgreSQL
* Docs: Swagger (via `swaggo/swag`)
* Container: Docker & Docker Compose

---

## 👨‍💻 Author

> Dibuat sebagai bagian dari penilaian teknikal test.

---
