# Task 3 — User API with Echo Framework (CRUD + Validation + Swagger)

This project is part of Sinau Koding Backend Bootcamp Batch 5.

It is a RESTful API for managing users, built using:
- Golang
- Echo Framework
- Validator (go-playground)
- Swagger (echo-swagger)

## 🚀 Features

✔ Get all users  
✔ Get user by ID  
✔ Create user  
✔ Update user  
✔ Delete user  
✔ Request body validation  
✔ Auto-increment user ID  

## 🧩 API Endpoints

| Method | Endpoint     | Description |
|-------|-------------|----------|
| GET   | /users       | Get all users |
| GET   | /users/{id}  | Get user by ID |
| POST  | /users       | Create new user |
| PUT   | /users/{id}  | Update user |
| DELETE | /users/{id} | Delete user |

---

## ⚙️ Run Application

Application runs at:
http://localhost:8080

Swagger Docs:
http://localhost:8080/swagger/index.html

---

## 📌 Example User Structure

{
"name": "John Doe",
"age": 22
}

Validation Rules:

- name → required  
- age → must be >= 0


