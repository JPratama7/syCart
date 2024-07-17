# SyCart
Simple implementation of a shopping cart using Golang

## Features
- Create a new cart
- Add items to the cart
- Remove items from the cart
- Get All Items in the cart
- Checkout the cart

## Database Design
![Database Design](diagram.png)

## Postman Collection
CLick Here to Download the Postman Collection  
[Postman Collection](Cart%20Spec.postman_collection.json)

## Installation
```bash
go mod tidy
go build -o main cmd/api/main.go
```

## Installation Using Docker
```bash
docker compose up -d
```

## Usage
```bash
./main
```