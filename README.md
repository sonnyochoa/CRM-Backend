# CRM-Backend
CRM Backend is built with Go. It is the server-side portion of a complete CRM application. Users are able to make HTTP requests to perform CRUD operations. With this project, I have implemented several different RESTful endpoints which allow you to perform CRUD operations.

CRM-Backend provides an API that includes functionality to read, write, update, and delete content. This application can easily be integrated with any front-end application. All requests are done through HTTP, and requests and responses are in JSON.

To test out the application you can use Postman or cURL.

## Overview
The application handles the following 5 operations for customers:

- Getting a single customer: `/customers/{id}`
- Getting all customers: `/customers`
- Creating a customer: `/customers`
- Updating a customer: `/customers/{id}`
- Deleting a customer: `/customers/{id}`

## Install
Clone the project to a local directory:
```bash
git clone https://github.com/sonnyochoa/CRM-Backend.git
```

Inside the project directory run:
```bash
go run main.go
```

## Examples

### Getting all customers: `/customers`    **GET**
#### EXAMPLE: `GET /customers`
**Output:**
```
{
    "1": {
        "id": 1,
        "name": "sonny",
        "role": "admin",
        "email": "admin@sonny.com",
        "phone": "1234567890"
    },
    "2": {
        "id": 2,
        "name": "panda",
        "role": "writer",
        "email": "writer@sonny.com",
        "phone": "5012346789"
    },
    "3": {
        "id": 3,
        "name": "bear",
        "role": "builder",
        "email": "builder@sonny.com",
        "phone": "3012456789",
        "contacted": true
    }
}
```

### Getting a single customer: `/customers/{id}`    **GET**
#### EXAMPLE: `GET /customers/1`
**Output:** 
```
{
    "id": 1,
    "name": "sonny",
    "role": "admin",
    "email": "admin@sonny.com",
    "phone": "1234567890"
}
```

### Creating a customer: `/customers`    **POST**
#### EXAMPLE: `POST /customers`
**Request Body Input:**
```
{
    "name": "22new customer",
    "role": "22test customer",
    "email": "22customer@test.com",
    "phone": "2225555555",
    "contacted": true
}
```
**Output:**
```
{
    "id": 200428821,
    "name": "22new customer",
    "role": "22test customer",
    "email": "22customer@test.com",
    "phone": "2225555555",
    "contacted": true
}
```

### Updating a customer: `/customers/{id}`    **PUT**
#### EXAMPLE: `PUT /customers/1`
**Request Body Input:**
```
{
    "name": "SONNY",
    "role": "SUPER ADMIN",
    "email": "SUPEADMIN@test.com",
    "phone": "9097061372",
    "contacted": true
}
```
**Output:**
```
{
    "id": 1,
    "name": "SONNY",
    "role": "SUPER ADMIN",
    "email": "SUPEADMIN@test.com",
    "phone": "9097061372",
    "contacted": true
}
```

### Deleting a customer: `/customers/{id}`    **DELETE**
#### EXAMPLE: `DELETE /customers/2`
**Output:**
```
{
    "1": {
        "id": 1,
        "name": "SONNY",
        "role": "SUPER ADMIN",
        "email": "SUPEADMIN@test.com",
        "phone": "9097061372",
        "contacted": true
    },
    "3": {
        "id": 3,
        "name": "bear",
        "role": "builder",
        "email": "builder@sonny.com",
        "phone": "3012456789",
        "contacted": true
    }
}
```