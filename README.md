# Microservices: Client and Server

This project consists of two services:

1. **Server**: Provides a get books API to return a list of books. Runs on port `8081`.
2. **Client**: Connects to the server to retrieve the book data and displays it. Runs on port `8080`.

Both services are written in Go and can be run using `go run main.go`.

## How to Run

Follow these steps to set up and run the microservices locally:

### 1. Clone the repository

Clone the repository and navigate to the project folder:

```
`git clone https://github.com/fadhlinw/learn-microservices.git`
```

setup your .env from example

### 2. Run Service Get Books (as a Server)

Services Provides backend services via gRPC, processes requests from clients, and returns responses, running on port 8081.
```
`cd learn-microservices/services` 
```
and next 
```
`go run main.go`
```
### 3. Run Client for 

Client Receive REST requests from users and forward them to the server using gRPC, running on port 8080.
```
`cd learn-microservices`
```
and next
```
`go run main.go`
```

