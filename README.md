# Golang Microservice

A simple yet powerful microservice architecture built using Golang. This project showcases a microservices setup with distinct services for users, products, and payments, demonstrating efficient data handling and communication between services using Kafka.

## Demo

Check out the demo video to see the microservice in action:
[Screencast from 31-03-24 01:42:31 AM IST.webm](https://github.com/SubhamMurarka/microService/assets/108292932/4e1e57eb-9f4c-4c4c-8bd2-3a0b1be26da3)

## Tech Stack

- **Language:** Go
- **Databases:**
  - MySQL (Users)
  - MongoDB (Products)
  - PostgreSQL (Payments)
- **Message Broker:** Apache Kafka (Product -> Payment)

## Services

### User Service

Handles user authentication and management. Users can sign up and log in. The user data is stored in a MySQL database.

### Product Service

Manages product data with full CRUD (Create, Read, Update, Delete) operations. Users can view a list of products. This service also acts as a Kafka producer, sending messages about product purchase. Product data is stored in a MongoDB database.

### Payment Service

Receives messages from the product service through Kafka. It processes the requests and stores payment data in a PostgreSQL database. This service acts as a Kafka consumer.

## Getting Started

### Prerequisites

Ensure you have the following installed:

- Git
- Docker
- Docker Compose
- Make

### Setup

1. **Clone the repository:**

    ```bash
    git clone https://github.com/ShubhamMurarka/microservice
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
- **Product Service:** Perform CRUD operations on products and view the list of products.
- **Payment Service:** Handle payment processing triggered by product updates.

This README provides a comprehensive overview of the Golang microservice project, including setup instructions, service descriptions, and usage guidelines. Enjoy working with this microservice architecture!
