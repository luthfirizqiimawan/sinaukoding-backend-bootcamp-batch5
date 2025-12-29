# 🧩 Task 3 — User API with Echo Framework (CRUD + Validation + Swagger)

Project ini merupakan pengembangan dari tugas sebelumnya dengan menambahkan:

- CRUD User API
- Request Body Validation (`validator.v10`)
- Auto Increment User ID
- Swagger API Documentation
- JSON Response Standar

Dibuat menggunakan **Golang + Echo Framework** sebagai bagian dari  
Bootcamp Sinau Koding — Backend Batch 5.

---

## 🛠 Tech Stack

- Golang
- Echo Web Framework
- Swagger (swaggo)
- Go Playground Validator v10

---

## 📁 Project Structure

```
task-3-user-api-echo-crud-validation-swagger
├── .gitignore
├── README.md
├── go.mod
├── go.sum
└── main.go
```

> Catatan: Folder `docs/` akan muncul setelah Swagger digenerate

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

| Method | Endpoint     | Keterangan       |
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
  "name": "John Doe",
  "age": 22
}
```

### ✅ Validation Rules

- `name` → required
- `age` → must be >= 0

---

## ✨ Notes

Generate ulang Swagger docs jika diperlukan:

```bash
swag init
```

---

## 🎯 Learning Objectives

- Implementasi REST API dengan Echo
- Validasi request body
- Swagger API Documentation
- Clean API structure & coding practice
