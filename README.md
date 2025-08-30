# vehicle-management-service

This repository contains the backend service for a Vehicle Management application.
The service provides a RESTful API to perform CRUD (Create, Read, Update, Delete) operations on vehicles,
as well as handle vehicle sales.
It is designed to be run in a containerized environment using Docker and Docker Compose.

## Features

- Create, Read, and Update vehicle information.
- Process vehicle sales, updating their status accordingly.
- List all vehicles with powerful filtering options (by brand, model, status).
- Sort vehicle lists it by price, year, or creation date.
- Paginated responses for efficient data retrieval on large datasets.

## Technologies

- **Containerization**: Docker and Docker Compose
- **Backend**: Go
- **Database**: PostgreSQL
- **API Documentation**: Swagger

## Getting Started

Follow these steps to get the application up and running on your local machine.

#### 1. Clone the repository:

```shell
git clone https://github.com/kalilventura/vehicle-management-service.git
cd vehicle-management-service
```

#### 2. Create the Environment File:

This project requires certain environment variables to be set.
You can find a template for these variables in the .env.example file.
To create your own .env file, run the following command:

```shell
cp .dev/.env.example .dev/.env
```

Then, edit the .env file to include the appropriate values for your setup.

#### 3. Building and running your application

When you're ready, start your application by running:

```shell
make dev-up
```

This command will build and start all the services defined in your docker-compose.yml file.

Your application will be accessible at `http://localhost:8080`, and the Swagger
documentation can be found at `http://localhost:8080/swagger/index.html`.

## API Documentation

The API provides several endpoints for managing vehicles. The base URL is `http://localhost:8080`.

For full, interactive documentation, you can likely access a Swagger UI instance at
`http://localhost:8080/swagger/index.html` or similar a path once the application is running.

## Contributing

Contributions are welcome! Please submit a pull request with a clear description of your changes.
