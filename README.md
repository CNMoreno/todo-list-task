# 📝 Todo List Task - Go API

This project is a **Go** API that allows managing a task list (**To-Do List**). The application uses **in-memory persistence** and handles authentication with **JWT**.

## 🚀 Features
- 📌 **User Management**: Create new users.
- 📌 **Task CRUD**: Create, retrieve, update, and delete tasks.
- 🔒 **JWT Authentication**: Token generation and validation.
- 🔄 **In-Memory Persistence**: Data is stored in memory while the API is running.
- ⚡ **Concurrent and Secure**: Uses `sync.RWMutex` to handle concurrent access.

## 📦 Installation

### 1️⃣ Clone the repository
```sh
git clone https://github.com/CNMoreno/todo-list-task.git
cd todo-list-task
```

### 2️⃣ Install dependencies
```sh
go mod tidy
```

### 3️⃣ Run the API
```sh
go run cmd/main.go
```

### 4️⃣ Test with `curl`
```sh
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "test", "password": "password"}'
```

## 🛠 Endpoints

### 🧑‍💻 Authentication & User Management
| Method | Endpoint  | Description               |
|--------|----------|---------------------------|
| POST   | `/users` | Creates a new user        |
| POST   | `/login` | Logs in and generates a JWT |

### ✅ Task Management
| Method | Endpoint      | Description               |
|--------|--------------|---------------------------|
| POST   | `/tasks`     | Creates a new task        |
| GET    | `/tasks`     | Retrieves all tasks       |
| GET    | `/tasks/:id` | Retrieves a specific task |
| PUT    | `/tasks/:id` | Updates a task            |
| DELETE | `/tasks/:id` | Deletes a task            |

## 🚀 Usage Examples

### 1️⃣ **Create a User**
```sh
curl -X POST http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"username": "test", "password": "password"}'
```

### 2️⃣ **Get Authentication Token**
```sh
curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{"username": "test", "password": "password"}'
```

### 3️⃣ **Create a Task**
```sh
curl -X POST http://localhost:8080/tasks \
     -H "Authorization: Bearer <TOKEN_HERE>" \
     -H "Content-Type: application/json" \
     -d '{"title": "Buy milk", "description":"I need milk for my coffee"}'
```

### 4️⃣ **Retrieve All Tasks**
```sh
curl -X GET http://localhost:8080/tasks \
     -H "Authorization: Bearer <TOKEN_HERE>"
```

## 🏗️ Architecture

📂 **Project Structure**
```
📂 todo-list-task
 ├── 📂 cmd                # Main entry point
 │    ├── main.go
 ├── 📂 internal
 │    ├── 📂 app            # Business logic
 │    ├── 📂 domain         # Business models
 │    ├── 📂 infrastructure # In-memory persistence and HTTP controllers
 │    ├── 📂 middleware     # Middleware logic
 │    ├── 📂 utils          # Utility functions
 ├── go.mod
 ├── go.sum
 ├── README.md
```

## 🛠 Technologies Used
- 🌐 **Go (Golang)**
- 🏗 **Gin (Web Framework)**
- 🔑 **JWT for authentication**
- 💾 **In-memory persistence using `sync.RWMutex`**

## 🏃 Testing
To run unit tests:
```sh
go test ./internal/infrastructure/http/  ./internal/utils
```
To generate coverage:
```sh
make coverage
```

## ⚖️ License
MIT 📄

---
✨ Contributions are welcome! If you have ideas to improve the project, feel free to open a PR or an issue. 🚀

