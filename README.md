# Go Http Template Server

A lightweight HTTP server built with Go that demonstrates routing, request validation, and HTML template rendering using only Go's standard library.

This project serves as an introduction to backend web development in Go and showcases how to render HTML pages dynamically using the `html/template` package.

---

## Features

* HTTP server powered by Go's `net/http` package
* Route handling with `http.ServeMux`
* HTML template rendering
* Request method validation
* Custom error handling
* Clean and beginner-friendly code structure
* No third-party dependencies

---

## Project Structure

```text
go-template-server/
|___ go.mod
├── main.go
└── templates/
    └── home.html
```

---

## Getting Started

### Prerequisites

* Go 1.20 or higher

Verify your installation:

```bash
go version
```

### Clone the Repository

```bash
git clone https://github.com/<chekus-dev>/go-template-server.git

cd go-template-server
```

### Run the Application

```bash
go run .
```

You should see:

```text
starting the server on port :7000
```

---

## Usage

Open your browser and visit:

```text
http://localhost:7000/home
```

The server will render the `home.html` template located in the `templates` directory.

---

## Route Reference

| Method | Endpoint | Description          |
| ------ | -------- | -------------------- |
| GET    | `/`      | Redirect to `/home`  |
| GET    | `/home`  | Render the Home page |

---

## Example Response

The application renders an HTML page similar to:

```html
<!DOCTYPE html>
<html>
<head>
    <title>HomePage</title>
</head>
<body style="background-color: lightgray;">
    <h1 style="color: blue;">
        Welcome To Chekus-dev little Go HTTP Server with HTML!
    </h1>
</body>
</html>
```

---

## Technologies Used

* Go
* net/http
* html/template

---

## Concepts Demonstrated

This project covers several important backend development concepts:

* Creating an HTTP server
* Defining route handlers
* Processing HTTP requests
* Validating request methods
* Rendering HTML templates
* Handling server-side errors
* Organizing a simple web application

---

## Future Enhancements

Potential improvements include:

* Multiple pages (Home, About, Projects, Contact)
* Shared template layouts
* Static CSS and JavaScript files
* Form handling
* Middleware support
* Custom 404 and 500 pages
* Logging and request monitoring

---

## Why This Project?

This project was built as a learning exercise to gain hands-on experience with Go's standard library and understand the fundamentals of web server development before moving on to larger frameworks and applications.

## Author
chekus-dev
