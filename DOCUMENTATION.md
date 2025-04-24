# GO REST API Documentation

## Installation & Run

```bash
cd backend
go mod tidy # install dependecies
go run main.go
```

### Base URL: `http://localhost:8080`

## Authentication

API uses JWT Authentication. Include the JWT token in the Authorization header for protected endpoints under `/api/`:

```
Authorization: Bearer <token>
```

## Endpoints

### Health Check

**GET** `/healthcheck`  
Check if the API is running and connected to the database.

**Response**: 200 OK

```json
{
  "message": "Hello, you have successfully connected to Postgres 🫶"
}
```

### Authentication

**POST** `/login`  
Authenticate a user and retrieve a JWT token.

**Request Body**:

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response**: 200 OK

```json
{
  "token": "jwt_token_here",
  "user_id": 1,
  "first_name": "John",
  "last_name": "Doe"
}
```

**Errors**:

- 400: Invalid request body
- 401: Invalid email or password

### Users

**GET** `/api/users`  
Get all users. Requires authentication.

**Response**: 200 OK

```json
[
  {
    "id": 1,
    "email": "user@example.com",
    "phone_number": "1234567890",
    "first_name": "John",
    "last_name": "Doe",
    "nick_name": "JD"
  }
]
```

**GET** `/api/user/{id}`  
Get user by ID. Requires authentication.

**Response**: 200 OK

```json
{
  "id": 1,
  "email": "user@example.com",
  "phone_number": "1234567890",
  "first_name": "John",
  "last_name": "Doe",
  "nick_name": "JD"
}
```

**Errors**:

- 400: Invalid user ID
- 500: Failed to retrieve user

### Items

**GET** `/api/items`  
Get all items. Requires authentication.

**Response**: 200 OK

```json
[
  {
    "id": 1,
    "name": "Lawn Mower",
    "description": "Gas-powered lawn mower in good condition",
    "category_id": 3,
    "owner_id": 1,
    "price": 1500,
    "date_listed": "2023-10-25T15:30:45Z",
    "quantity": 1,
    "available": true
  }
]
```

**POST** `/api/items`  
Create a new item. Requires authentication.

**Request Body**:

```json
{
  "name": "Lawn Mower",
  "description": "Gas-powered lawn mower in good condition",
  "category_id": 3,
  "price": 1500,
  "quantity": 1,
  "available": true
}
```

**Response**: 200 OK

```json
{
  "id": 1,
  "name": "Lawn Mower",
  "description": "Gas-powered lawn mower in good condition",
  "category_id": 3,
  "owner_id": 1,
  "price": 1500,
  "date_listed": "2023-10-25T15:30:45Z",
  "quantity": 1,
  "available": true
}
```

**Errors**:

- 400: Invalid request body
- 500: Failed to create item

**GET** `/api/items/available`  
Get available items for rent with owner information. Requires authentication.

**Response**: 200 OK

```json
[
  {
    "id": 1,
    "name": "Lawn Mower",
    "description": "Gas-powered lawn mower in good condition",
    "price_per_day": 15.0,
    "owner_id": 2,
    "owner_name": "Jane Smith",
    "available": true
  }
]
```

### Categories

**GET** `/api/categories`  
Get all categories. Requires authentication.

**Response**: 200 OK

```json
[
  {
    "id": 1,
    "name": "Tools",
    "description": "Hand and power tools for various projects"
  }
]
```

### Rentals

**POST** `/api/rentals`  
Create a rental request. Requires authentication.

**Request Body**:

```json
{
  "item_id": 1,
  "start_date": "2023-10-30T10:00:00Z",
  "end_date": "2023-10-31T10:00:00Z",
  "total_price": 1500
}
```

**Response**: 200 OK

```json
{
  "id": 1,
  "item_id": 1,
  "renter_id": 2,
  "start_date": "2023-10-30T10:00:00Z",
  "end_date": "2023-10-31T10:00:00Z",
  "status": "pending",
  "total_price": 1500
}
```

**Errors**:

- 400: Invalid request
- 500: Failed to create rental request

## Status Codes

- 200: Success
- 400: Bad Request - Invalid input or parameters
- 401: Unauthorized - Authentication required
- 403: Forbidden - Insufficient permissions
- 404: Not Found - Resource doesn't exist
- 500: Internal Server Error - Server-side problem

## Data Models

Check `models/`
