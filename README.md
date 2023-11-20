# Distributedsystemsproject

Overview

The application provides a RESTful API for basic CRUD operations on a collection of items. It is structured as a Go program and is divided into several packages: database, helper, and otel. The database package handles interactions with the underlying data store, helper provides utility functions, and otel integrates OpenTelemetry for distributed tracing.
Dependencies

The following external packages are imported and utilized in the application:

    github.com/gin-gonic/gin: The Gin web framework for handling HTTP requests.
    go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: OpenTelemetry instrumentation for HTTP servers.
    go.opentelemetry.io/otel/trace: OpenTelemetry tracing package.

OpenTelemetry Setup

The application sets up OpenTelemetry to enable distributed tracing. It initializes the OpenTelemetry SDK with the service name "Todo" and version "0.1". The tracing is automatically wrapped around the HTTP server using the gin.WrapH middleware.
Database Initialization

The database package is initialized using database.Init(), ensuring that the application is ready to interact with the underlying data store.
API Endpoints
GET /test

    Returns a JSON response containing a randomly generated ID with a length of 32 characters.

GET /

    Returns a JSON response indicating that the API is working.

GET /posts

    Returns a JSON response containing a list of all items in the database, each represented in a shortened format.

GET /post/:id

    Returns a JSON response containing details of the item with the specified ID.

POST /post

    Adds a new item to the database based on the JSON payload received in the request.

PUT /post

    Updates an existing item in the database. Expects the ID of the item to be updated in the form data.

DELETE /post/:id

    Deletes the item with the specified ID from the database.

Run the Application

The application listens on port 9988. To start the server, run the main function in the main package. The server will be accessible at http://localhost:9988.

------------------------
generated by ChatGPT
