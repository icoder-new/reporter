# Financial Report Project

This project is a financial reporting system implemented in Go (version 1.20+), utilizing the following libraries: gin-gonic, gORM (psql), golang-jwt, godotenv, and viper for configuration.

## Routes

The project includes the following routes:

### `/`

- **Method:** GET
- **Description:** Ping route to check server availability.
- **Response:**
  - Status: 200 OK
  - Body:
    ```json
    {
      "message": "ping - pong"
    }
    ```

### `/api/v1/auth/register`

- **Method:** POST
- **Description:** User registration route.
- **Request Body:**
  - Content-Type: application/json
  - Body:
    ```json
    {
      "firstname": "first_name",
      "lastname": "last_name",
      "username": "example_user",
      "email": "example_user@domain.com",
      "password": "password123"
    }
    ```
- **Response:**
  - Status: 200 OK (on successful registration)
  - Status: 400 Bad Request (if the request body is invalid or registration fails)
  - Body (on successful registration):
    ```json
    {
      "message": "success",
      "id": "will be returned ID"
    }
    ```

### `/api/v1/auth/login`

- **Method:** POST
- **Description:** User login route.
- **Request Body:**
  - Content-Type: application/json
  - Body:
    ```json
    {
      "username": "example_user",
      "email": "example_user@domain.com",
      "password": "password123"
    }
    ```
- **Response:**
  - Status: 200 OK (on successful login)
  - Status: 400 Bad Request (if the request body is invalid or login fails)
  - Body (on successful login):
    ```json
    {
      "token": "your_jwt_token"
    }
    ```

### `/api/v1/user`

- **Method:** GET
- **Description:** Get user information route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **Response:**
  - Status: 200 OK (on successful retrieval)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful retrieval):
    ```json
    {
      "firstname": "first_name",
      "lastname": "last_name",
      "username": "example_user",
      "email": "example_user@domain.com",
      "is_active": true
      // "picture": "path to the photo profile"
    }
    ```

### `/api/v1/user`

- **Method:** PUT
- **Description:** Update user information route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **Request Body:**
  - Content-Type: application/json
  - Body:
    ```json
    {
      "firstname": "first_name",
      "lastname": "last_name",
      "username": "example_user",
      "email": "example_user@domain.com",
      "password": "password123"
    }
    ```
- **Response:**
  - Status: 200 OK (on successful update)
  - Status: 400 Bad Request (if the request body is invalid or update fails)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful update):
    ```json
    {
      "message": "success",
      "data": [
        {/* changed user info */}
      ]
    }
    ```

### `/api/v1/user/restore`

- **Method:** GET
- **Description:** Restore deleted user route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **Response:**
  - Status: 200 OK (on successful restore)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful restore):
    ```json
    {
      "message": "success"
    }
    ```

### `/api/v1/user/delete`

- **Method:** DELETE
- **

Description:** Delete user route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **Response:**
  - Status: 200 OK (on successful deletion)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful deletion):
    ```json
    {
      "message": "success"
    }
    ```

### `/api/v1/user/change`

- **Method:** PATCH
- **Description:** Change user profile picture route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **Request Body:**
  - Content-Type: multipart/form-data
  - Form Data:
    - `picture`: <file>
- **Response:**
  - Status: 200 OK (on successful picture change)
  - Status: 400 Bad Request (if the request body is invalid or change fails)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful picture change):
    ```json
    {
      "message": "success",
      "data": [
        {
            // ...
            "path": "file_name.jpg"
            //...
        }
      ]
    }
    ```

### `/api/v1/user/upload`

- **Method:** PATCH
- **Description:** Upload user profile picture route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **Request Body:**
  - Content-Type: multipart/form-data
  - Form Data:
    - `picture`: <file>
- **Response:**
  - Status: 200 OK (on successful picture upload)
  - Status: 400 Bad Request (if the request body is invalid or upload fails)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful picture upload):
    ```json
    {
      "message": "success"
    }
    ```

### `/api/v1/account`

- **Method:** GET
- **Description:** Get all accounts route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **Response:**
  - Status: 200 OK (on successful retrieval)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful retrieval):
    ```json
    [
      {
        "id": 1,
        "user_id": 3,
        "name": "Savings Account",
        "balance": 1000.0
      },
      {
        "id": 2,
        "user_id": 2,
        "name": "Expense Account",
        "balance": -500.0
      }
    ]
    ```

### `/api/v1/account/:id`

- **Method:** GET
- **Description:** Get account information route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **URL Parameters:**
  - `id`: Account ID
- **Response:**
  - Status: 200 OK (on successful retrieval)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful retrieval):
    ```json
    {
      "id": 2,
      "user_id": 8,
      "name": "Savings Account",
      "balance": 1000.0
    }
    ```

### `/api/v1/account`

- **Method:** POST
- **Description:** Create a new account route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **Request Body:**
  - Content-Type: application/json
  - Body:
    ```json
    {
      "name": "New Account",
      "balance": 0.0
    }
    ```
- **Response:**
  - Status: 201 Created (on successful creation)
  - Status

: 400 Bad Request (if the request body is invalid or creation fails)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful creation):
    ```json
    {
      "id": 3,
      "user_id": 6,
      "name": "New Account",
      "balance": 0.0
    }
    ```

### `/api/v1/account/:id`

- **Method:** PUT
- **Description:** Update account information route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **URL Parameters:**
  - `id`: Account ID
- **Request Body:**
  - Content-Type: application/json
  - Body:
    ```json
    {
      "name": "Updated Account",
      "balance": 500.0
    }
    ```
- **Response:**
  - Status: 200 OK (on successful update)
  - Status: 400 Bad Request (if the request body is invalid or update fails)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful update):
    ```json
    {
      "message": "success",
      "data": [{
        // updated data
      }]
    }
    ```

### `/api/v1/account/:id/restore`

- **Method:** GET
- **Description:** Restore deleted account route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **URL Parameters:**
  - `id`: Account ID
- **Response:**
  - Status: 200 OK (on successful restore)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful restore):
    ```json
    {
      "message": "success"
    }
    ```

### `/api/v1/account/:id`

- **Method:** DELETE
- **Description:** Delete account route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **URL Parameters:**
  - `id`: Account ID
- **Response:**
  - Status: 200 OK (on successful deletion)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful deletion):
    ```json
    {
      "message": "success"
    }
    ```

### `/api/v1/account/:id/change`

- **Method:** PATCH
- **Description:** Change account picture route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **URL Parameters:**
  - `id`: Account ID
- **Request Body:**
  - Content-Type: multipart/form-data
  - Form Data:
    - `picture`: <file>
- **Response:**
  - Status: 200 OK (on successful picture change)
  - Status: 400 Bad Request (if the request body is invalid or change fails)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful picture change):
    ```json
    {
      "message": "success",
      "data": [{
        // ...
        "path": "file_name.jpg"
        // ...
      }]
    }
    ```

### `/api/v1/account/:id/upload`

- **Method:** PATCH
- **Description:** Upload account picture route.
- **Headers:**
  - Authorization: Bearer <JWT_TOKEN>
- **URL Parameters:**
  - `id`: Account ID
- **Request Body:**
  - Content-Type: multipart/form-data
  - Form Data:
    - `picture`: <file>
- **Response:**
  - Status: 200 OK (on successful picture upload)
  - Status: 400 Bad Request (if the request body is invalid or

 upload fails)
  - Status: 401 Unauthorized (if the JWT token is missing or invalid)
  - Body (on successful picture upload):
    ```json
    {
      "message": "success",
      "data": [{
        // ...
        "path": "file_name.jpeg"
        // ...
      }]
    }
    ```

## Database Models

The project utilizes the following database models:

### User

- `id` (integer): Unique identifier for the user.
- `fistname` (string): User's first name.
- `lastname` (string): User's last name.
- `username` (string): User's username   | unique.
- `email` (string): User's email address | unique.
- `password` (string): User's hashed password.
- `is_active` (bool): for soft deleting
- `deleted_at` (datetime): Date and time when the user was deleted.
- `created_at` (datetime): Date and time when the user was created.
- `updated_at` (datetime): Date and time when the user was last updated.

### Account

- `id` (integer): Unique identifier for the account.
- `user_id` (integer): ID of the user who owns the account.
- `name` (string): Account name.
- `balance` (decimal): Account balance.
- `is_active` (bool): for soft deleting
- `deleted_at` (datetime): Date and time when the account was deleted.
- `created_at` (datetime): Date and time when the account was created.
- `updated_at` (datetime): Date and time when the account was last updated.

## Authentication and Authorization

The project utilizes JWT (JSON Web Tokens) for authentication and authorization. When a user registers or logs in, a JWT token is generated and returned. This token should be included in the `Authorization` header for authenticated routes as `Bearer <JWT_TOKEN>`.

Authenticated routes require a valid JWT token for access. If the token is missing, expired, or invalid, the server will respond with a `401 Unauthorized` status.

## Running the Application

To run the application, follow these steps:

1. Set the required environment variables.
2. Install the project dependencies using `go mod tidy`.
3. Build the project using `go build -o reporter`.
4. Run the executable file `reporter`.

The application should now be running on the specified port, and you can access the routes using a tool like Postman or cURL.