# Banking API

## Overview

This project is a Banking API implementation using the Hexagonal Architecture (also known as Ports and Adapters) pattern. It is part of a larger banking microservices ecosystem developed in Go (Golang). The goal of this API is to provide a robust and flexible interface for managing banking operations, ensuring separation of concerns and easier testing.

## Architecture

The application follows the Hexagonal Architecture pattern, which promotes:

- **Separation of Concerns**: Core business logic is isolated from external dependencies.
- **Testability**: Each component can be tested independently.
- **Flexibility**: Easily swap out components (e.g., databases or external services) without affecting the core logic.

## Installation

To set up the project locally, follow these steps:

1. Clone the repository:
git clone https://github.com/yourusername/banking-api.git
cd banking-api

2. Install dependencies

3. Set up environment variables as needed (e.g., database connection strings).

## Usage
To run the application locally:
SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWD=root DB_ADDR=localhost DB_PORT=3307 DB_NAME=banking go run main.go

