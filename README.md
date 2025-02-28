# Golang Basic Web Server

A simple web server built using Golang that serves static files and handles HTTP requests.

## Features
- Serves static files from the `./static` directory.
- Handles `GET` requests on `/hello`.
- Handles `POST` requests on `/form` to process form data.
- Runs on port `8080`.

## Installation
Ensure you have **Go** installed on your system. You can download it from [golang.org](https://golang.org/dl/).

Clone the repository:
```sh
git clone https://github.com/vineet12344/GO---Basic-Web-Server-.git
cd GO---Basic-Web-Server-

Usage

Run the server using:

go run main.go

The server will start on http://localhost:8080.
Endpoints
GET /hello

Responds with:

Hello Ebery Nyan

POST /form

Accepts form data (name, address) and returns a response with submitted values:

POST request successful

Name: <your_name>
Address: <your_address>

Project Structure

📦 GO---Basic-Web-Server-
├── 📂 static         # Static files (HTML, CSS, JS, etc.)
├── 📄 main.go        # Main server code
└── 📄 README.md      # Project documentation

License

This project is open-source and available under the MIT License.
Author

Vineet Salve
