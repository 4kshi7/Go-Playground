## 📘 Go CRUD API with Gin

This is a simple RESTful API built using [Go](https://golang.org/) and the [Gin](https://github.com/gin-gonic/gin) web framework. It demonstrates basic CRUD operations (Create, Read, Update, Delete) using in-memory storage.

## 🚀 Features

- GET all items
- GET a single item by ID
- POST a new item
- PUT (update) an item
- DELETE an item

## 📦 Dependencies

- [Gin](https://github.com/gin-gonic/gin): HTTP web framework for Go

## 📁 Project Structure

```bash
go-crud-api/
├── go.mod
├── go.sum
└── main.go
````

## 🔧 Installation

1. **Clone the repo:**

   ```bash
   git clone https://github.com/4kshi7/go-playground.git
   cd go-playground
   ```

2. **Initialize Go module (if not done):**

   ```bash
   go mod init go-crud-api
   ```

3. **Install dependencies:**

   ```bash
   go get github.com/gin-gonic/gin
   ```

4. **Run the server:**

   ```bash
   go run main.go
   ```

The server will start at [http://localhost:8080](http://localhost:8080)

## 📬 API Endpoints

| Method | Endpoint     | Description             |
| ------ | ------------ | ----------------------- |
| GET    | `/books`     | Get all books           |
| GET    | `/books/:id` | Get book by ID          |
| POST   | `/books`     | Create a new book       |
| PUT    | `/books/:id` | Update an existing book |
| DELETE | `/books/:id` | Delete a book           |

## 📄 Sample `Book` JSON

```json
{
  "id": "1",
  "title": "The Go Programming Language",
  "author": "Alan Donovan"
}
```

## 📚 Learn More

* [Go Documentation](https://golang.org/doc/)
* [Gin Web Framework](https://gin-gonic.com/)

## 🛠️ Future Improvements

* Database integration (PostgreSQL, MySQL)
* Middleware for logging and authentication
* Request validation
* Unit tests

---

