# Go HTTP Server Example

This is a simple HTTP server written in Go using the standard `net/http` package.  
It demonstrates how to:

- Serve static files
- Handle form submissions
- Create basic HTTP endpoints

The project is based on this tutorial:  
https://www.youtube.com/watch?v=jFfo23yIWac

---

## 🚀 Features

- Serves static files from the `./static` directory
- Handles POST requests from an HTML form at `/form`
- Exposes a simple GET endpoint at `/hello`
- Uses only Go standard library (no external dependencies)

---

## 📁 Project Structure

```text
.
├── main.go
└── static/
    └── index.html
````

> The `static` directory should contain your HTML files (for example, a form).

---

## 🛠 Requirements

* Go 1.16 or later

Check your Go version:

```bash
go version
```

---

## ▶️ How to Run

1. Clone the repository:

```bash
git clone <your-repo-url>
cd <project-folder>
```

2. Run the server:

```bash
go run main.go
```

3. Open your browser and visit:

* [http://localhost:8080/](http://localhost:8080/) → static files
* [http://localhost:8080/hello](http://localhost:8080/hello) → simple GET endpoint

---

## 📤 Form Endpoint

### `POST /form`

This endpoint processes form submissions.

Expected form fields:

* `name`
* `address`

Example response:

```text
POST request successful
Name = John
Address = Main Street
```

---

## 📌 Notes

* The server listens on port `8080`
* Requests to unsupported paths or methods return a `404` error
* This project is intended for learning and experimentation

---

## 📄 License

This project is provided for educational purposes.