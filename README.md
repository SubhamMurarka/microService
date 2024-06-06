# Golang Microservice

A simple yet powerful microservice architecture built using Golang. This project showcases a microservices setup with distinct services for users, products, and payments, demonstrating efficient data handling and communication between services using Kafka.

## Demo

Check out the demo video to see the microservice in action:
[Screencast from 31-03-24 01:42:31 AM IST.webm](https://github.com/SubhamMurarka/microService/assets/108292932/4e1e57eb-9f4c-4c4c-8bd2-3a0b1be26da3)

## Tech Stack

### Language
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

The entire microservice architecture is built using Go, a statically typed, compiled programming language designed for simplicity and efficiency.

### Framework
![Fiber](https://img.shields.io/badge/Fiber-00ADD8?style=for-the-badge&logo=go-fiber&logoColor=white)

All microservices use the Fiber framework, an Express-inspired web framework built on Fasthttp, for handling HTTP requests efficiently.

### Databases

![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white)

- **MySQL**: Used by the User Service to store user data.
- **MongoDB**: Used by the Product Service to manage product data with full CRUD operations.
- **PostgreSQL**: Used by the Payment Service to store payment data.

### Message Broker

![Kafka](https://img.shields.io/badge/Apache%20Kafka-231F20?style=for-the-badge&logo=apache-kafka&logoColor=white)

Apache Kafka is utilized for communication between the Product Service and the Payment Service, facilitating efficient message passing.

### Containerization and Orchestration

![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![Docker Compose](https://img.shields.io/badge/Docker_Compose-2496ED?style=for-the-badge&logo=docker&logoColor=white)

Docker is used to containerize the microservices, and Docker Compose orchestrates the multi-container Docker applications, ensuring all services, databases, and Kafka run seamlessly.

### Build Tool

![Make](https://img.shields.io/badge/Make-3776AB?style=for-the-badge&logo=gnu-make&logoColor=white)

Make is used to automate the setup and teardown of the Docker containers, simplifying the process of managing the development environment.

## Services

### User Service

Handles user authentication and management. Users can sign up and log in. The user data is stored in a MySQL database.

### Product Service

Manages product data with full CRUD (Create, Read, Update, Delete) operations. Users can view a list of products. This service also acts as a Kafka producer, sending messages about product purchases. Product data is stored in a MongoDB database.

### Payment Service

Receives messages from the product service through Kafka. It processes the requests and stores payment data in a PostgreSQL database. This service acts as a Kafka consumer.

## Getting Started

### Prerequisites

Ensure you have the following installed:

- ![Git](https://img.shields.io/badge/Git-F05032?style=for-the-badge&logo=git&logoColor=white)
- ![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
- ![Docker Compose](https://img.shields.io/badge/Docker_Compose-2496ED?style=for-the-badge&logo=docker&logoColor=white)
- ![Make](https://img.shields.io/badge/Make-3776AB?style=for-the-badge&logo=gnu-make&logoColor=white)

### Setup

1. **Clone the repository:**

    ```bash
    git clone https://github.com/SubhamMurarka/microService.git
    cd microservice
    ```

2. **Start the project:**

    ```bash
    make up
    ```

    This command will start all the necessary containers for the microservices, databases, and Kafka.

3. **Stop the project:**

    ```bash
    make down
    ```

    This command will stop all running containers.

## Usage

Once the services are up and running, you can interact with them through their respective endpoints.

- **User Service:** Manage user signup and login.
- **Product Service:** Perform CRUD operations on products and view the list of products and buy products as well.
- **Payment Service:** Handle payment processing triggered by product purchase.

This README provides a comprehensive overview of the Golang microservice project, including setup instructions, service descriptions, and usage guidelines. Enjoy working with this microservice architecture!
