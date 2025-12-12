# Go Movies REST API

This project is a simple **RESTful API written in Go** using the **Gorilla Mux** router.  
It demonstrates basic CRUD operations (Create, Read, Update, Delete) over an in-memory movie collection.

This API is intended for **learning purposes** and does not use a database.

---

## 🚀 Features

- REST API built with Go and `net/http`
- Routing with `github.com/gorilla/mux`
- CRUD operations for movies
- JSON request and response handling
- In-memory data storage

---

## 📁 Project Structure

```text
.
├── go.mod
├── go.sum
└── main.go
````

---

## 🛠 Requirements

* Go 1.16 or later

Check your Go version:

```bash
go version
```

---

## 📦 What is Gorilla Mux?

[Gorilla Mux](https://github.com/gorilla/mux) is a powerful and widely used HTTP router and URL matcher for Go.

It extends Go’s standard `net/http` package by providing:

- Flexible request routing
- URL parameters (path variables)
- Method-based routing (GET, POST, PUT, DELETE, etc.)
- Middleware support
- Better control over request handling

In this project, Gorilla Mux is used to:

- Define RESTful routes such as `/movies/{id}`
- Extract URL parameters using `mux.Vars`
- Restrict handlers to specific HTTP methods

Gorilla Mux integrates seamlessly with Go’s standard HTTP server and is commonly used for building REST APIs and microservices.

---

## ▶️ How to Run

1. Clone the repository:

```bash
git clone <your-repo-url>
cd <project-folder>
```

2. Download dependencies:

```bash
go mod tidy
```

3. Run the server:

```bash
go run main.go
```

The server will start at:

```
http://localhost:8000
```

---

## 📚 API Endpoints

### Get all movies

**GET** `/movies`

Response:

```json
[
  {
    "id": "1",
    "isbn": "438227",
    "title": "Star Wars",
    "director": {
      "firstName": "George",
      "lastName": "Lucas"
      }
  }
]
```

---

### Get a movie by ID

**GET** `/movies/{id}`

---

### Create a new movie

**POST** `/movies`

Request body:

```json
{
  "isbn": "999999",
  "title": "Interstellar",
  "director": {
    "firstName": "Christopher",
    "lastName": "Nolan"
  }
}
```

---

### Update a movie

**PUT** `/movies/{id}`

Request body:

```json
{
  "isbn": "999999",
  "title": "Interstellar (Updated)",
  "director": {
    "firstName": "Christopher",
    "lastName": "Nolan"
  }
}
```

---

### Delete a movie

**DELETE** `/movies/{id}`

---

## ⚠️ Important Notes

* Data is stored **in memory**, so all movies are lost when the server restarts
* IDs are generated randomly for new movies
* This project is not production-ready
* No authentication or validation is implemented

---

## 🧪 Testing with curl

Example:

```bash
curl http://localhost:8000/movies
```

---

## 📌 Future Improvements

* Add persistent storage (database)
* Add request validation
* Add error handling and HTTP status codes
* Add authentication
* Use environment variables for configuration

---

## 📄 License

This project is provided for educational purposes.