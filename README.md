Restaurant Management Backend

This project is a backend service for managing a restaurant. It is built using Go (Golang) and provides a set of RESTful APIs to handle various operations such as managing food items, orders, invoices, and authentication.

Features


User authentication and authorization
Manage food items (CRUD operations)
Handle customer orders and order items
Generate and manage invoices
Integration with MongoDB for data storage

Technologies Used

Programming Language: Go (Golang)
Framework: Gin for HTTP routing and middleware
Database: MongoDB
Authentication: JWT (JSON Web Tokens)

Getting Started

Prerequisites

Go 1.16 or higher
MongoDB
Installation

Clone the repository:

sh git clone cd

Install dependencies:

sh go mod download

Set up environment variables:

Create a .env file in the root directory and add the following variables:

env MONGO_URI= JWT_SECRET=

Run the application:

sh go run main.go

API Endpoints

The following are some of the main API endpoints provided by the backend:

Authentication

POST /login - Authenticate a user and get a JWT token
POST /register - Register a new user
Food Items

GET /foods - Get a list of all food items
POST /foods - Create a new food item
GET /foods/:id - Get details of a specific food item
PUT /foods/:id - Update a food item
DELETE /foods/:id - Delete a food item
Orders

GET /orders - Get a list of all orders
POST /orders - Create a new order
GET /orders/:id - Get details of a specific order
Invoices

GET /invoices - Get a list of all invoices
POST /invoices - Create a new invoice
GET /invoices/:id - Get details of a specific invoice
