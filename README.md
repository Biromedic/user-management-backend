# Go Programming Exercises

This repository contains solutions to four different Go programming questions (Q1, Q2, Q3, and Q4). Each question is implemented in its respective directory and demonstrates various aspects of Go programming, including sorting, recursion, data manipulation, and building a RESTful API with middleware and Swagger documentation.

## Table of Contents

1. [Q1: Sorting Words by 'a' Count](#q1-sorting-words-by-a-count)
2. [Q2: Recursive Function](#q2-recursive-function)
3. [Q3: Most Repeated Item in a Slice](#q3-most-repeated-item-in-a-slice)
4. [Q4: RESTful API with Swagger Documentation](#q4-restful-api-with-swagger-documentation)
---

## Q1: Sorting Words by 'a' Count

### Description

* This program sorts a slice of strings based on the number of occurrences of the letter 'a' (case-insensitive) in each word. If two words have the same number of 'a's, the longer word comes first.

* The program will print the sorted list of words.

## Q2: Recursive Function

* This program demonstrates a simple recursive function that prints values based on the current iteration. The function stops when the current value exceeds 3.

```plain
The program will print the following sequence:
2
4
9
```

## Q3: Most Repeated Item in a Slice

* This program finds the most repeated item in a slice of strings and prints it. It also counts the occurrences of each item in the slice.

* The program will print the most repeated item in the slice.

## Q4: RESTful API with Swagger Documentation

- This project implements a RESTful API for managing users. It includes CRUD operations, middleware for logging and CORS, and Swagger documentation for API endpoints.

### Features

- CRUD Operations: Create, Read, Update, and Delete users.
- Middleware: Logging and CORS middleware.
- Swagger Documentation: Automatically generated API documentation using Swagger.

### Files

- Q4/go.mod: Go module file with dependencies.
- Q4/main.go: Main program file to start the server.
- Q4/config/cors.go: CORS middleware implementation.
- Q4/docs/: Swagger documentation files.
- Q4/internal/database/connection.go: Database connection setup.
- Q4/internal/handler/user_handlers.go: HTTP handlers for user operations.
- Q4/internal/helpers/error_handlers.go: Error handling utilities.
- Q4/internal/middleware/logging_middleware.go: Logging middleware.
- Q4/internal/model/user.go: User model definition.
- Q4/internal/repository/: Repository layer for database operations.
- Q4/internal/routes/routes.go: API route setup.
- Q4/internal/service/user_service.go: Service layer for user operations.
- Q4/tests/: Unit and integration tests.


> Access the Swagger documentation at http://localhost:8080/swagger/.

### API Endpoints

- GET /users: Get all users.
- GET /users/{id}: Get a user by ID.
- POST /users: Create a new user.
- PUT /users/{id}: Update a user by ID.
- DELETE /users/{id}: Delete a user by ID.

### Tests

- The Q4 project includes both unit tests and integration tests to ensure the correctness of the code. The tests are located in the Q4/tests/ directory.

  -   Unit Tests
  -  Integration Tests

        - Unit tests are written for the service layer to test individual functions in isolation. Mock repositories are used to simulate database interactions.
        - Integration tests are written for the HTTP handlers to test the API endpoints. These tests simulate HTTP requests and verify the responses.

```plain
Test Cases(Unit Tests):

TestUserService_GetAllUsers: Tests the GetAllUsers function.

TestUserService_CreateUser: Tests the CreateUser function.

TestUserService_GetUserByID: Tests the GetUserByID function.

TestUserService_DeleteUser: Tests the DeleteUser function.


Test Cases(Integration Tests):

TestUserHandler_GetAllUsers: Tests the GET /users endpoint.

TestUserHandler_GetUserByID_ValidID: Tests the GET /users/{id} endpoint with a valid ID.

TestUserHandler_GetUserByID_NotFound: Tests the GET /users/{id} endpoint with an invalid ID.

TestUserHandler_CreateUser: Tests the POST /users endpoint.

TestUserHandler_UpdateUser: Tests the PUT /users/{id} endpoint.

TestUserHandler_DeleteUser_ValidID: Tests the DELETE /users/{id} endpoint with a valid ID.

TestUserHandler_DeleteUser_NotFound: Tests the DELETE /users/{id} endpoint with an invalid ID.
```