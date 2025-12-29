## 📌 User CRUD REST API — Echo + Swagger + Validation

Project ini merupakan implementasi REST API sederhana menggunakan **Golang + Echo Framework** dengan fitur:

- CRUD User  
- Validation (`go-playground/validator`)
- Auto Increment ID  
- Swagger API Documentation  
- JSON Response Standardized  

### 🛠 Teknologi yang digunakan

- Golang  
- Echo Framework  
- Swagger Docs  
- Validator v10  

---

## 📁 API Endpoints

| Method | Endpoint     | Description     |
|--------|-------------|----------------|
| GET    | /users      | Get all users  |
| GET    | /users/{id} | Get user by ID |
| POST   | /users      | Create user    |
| PUT    | /users/{id} | Update user    |
| DELETE | /users/{id} | Delete user    |

---

## ⚙️ Run Application

Jalankan aplikasi:

```bash
go run main.go
```

Aplikasi berjalan di:

```
http://localhost:8080
```

Swagger Docs:

```
http://localhost:8080/swagger/index.html
```

---

## 🧩 Example User Payload

```json
{
  "name": "John Doe",
  "age": 22
}
```

### Validation Rules

- name → required  
- age → must be ≥ 0  

---

## 🧪 Sample Responses

### ✅ Success (Create User)

```json
{
  "id": 4,
  "name": "John Doe",
  "age": 22
}
```

### ⚠️ Validation Error

```json
{
  "error": "Invalid input"
}
```

### ❌ User Not Found

```json
{
  "error": "User not found"
}
```

---

## ✨ Key Learning Points

- Implementasi REST API menggunakan Echo  
- Penyimpanan data dalam slice (in-memory)  
- Auto-increment user ID  
- Input validation menggunakan validator.v10  
- Error handling best practice  
- Swagger API Documentation  
