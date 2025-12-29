# 🧩 Go Echo User API — CRUD + Swagger + Validation

Project ini merupakan implementasi REST API sederhana menggunakan **Golang & Echo Framework** dengan fitur:

- CRUD User (Create, Read, Update, Delete)
- Validation menggunakan `validator.v10`
- Swagger API Documentation
- In-memory data (slice)
- JSON response terstruktur

Swagger documentation berada pada folder:

```
/docs
```

---

## 🛠 Tech Stack

- Golang
- Echo Framework
- Swagger (swaggo)
- Validator v10

---

## 📁 Project Structure

```
.
├── docs/           # Swagger generated files
├── .env
├── go.mod
├── go.sum
└── main.go
```

---

## 🚀 Run Application

Jalankan aplikasi:

```bash
go run main.go
```

Aplikasi berjalan pada:

```
http://localhost:8080
```

Swagger Docs:

```
http://localhost:8080/swagger/index.html
```

---

## 📌 API Endpoints

| Method | Endpoint     | Deskripsi        |
|--------|-------------|-----------------|
| GET    | /users      | Get all users   |
| GET    | /users/{id} | Get user by ID  |
| POST   | /users      | Create user     |
| PUT    | /users/{id} | Update user     |
| DELETE | /users/{id} | Delete user     |

---

## 🧩 Example Request Body

```json
{
  "id": 4,
  "name": "John Doe",
  "age": 22
}
```

### ✅ Validation Rules

- `id` → required
- `name` → required
- `age` → required & must be ≥ 0

---

## ✨ Notes

Untuk generate ulang Swagger docs:

```bash
swag init
```

Folder `docs/` akan diperbarui otomatis.

---

## 🎯 Learning Focus

- REST API dengan Echo
- Middleware & Validation
- Swagger Documentation
- Basic API Best Practices

---

💡 Project ini dibuat sebagai bagian dari  
**Bootcamp Sinau Koding — Backend (Golang & PostgreSQL)**
