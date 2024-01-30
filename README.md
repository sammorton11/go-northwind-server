# Go Northwind Server with Docker

This Go project is a simple HTTP server that interacts with the Northwind SQLite 3 database. The server is Dockerized for easy deployment.
Prerequisites

Before you begin, ensure that you have Docker installed on your machine.
Getting Started

## Clone the repository:

    git clone https://github.com/your-username/go-northwind-server.git


## Navigate to the project directory:

bash

cd go-northwind-server


## Build and run the Docker container:

    docker build -t go-northwind-server .
    docker run -p 8080:8080 go-northwind-server

The server will be running at http://localhost:8080.


## Available Endpoints

    /customers: Get customers and their orders from the Northwind database
    /orders: Get all orders from the Northwind database

## Stop the Docker Container

To stop the Docker container, find its ID using:


docker ps


## Then stop the container:

docker stop <container-id>

Replace <container-id> with the actual ID of your running container.
