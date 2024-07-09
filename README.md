<h1 style="text-align: center;">
    Notifications Management System
</h1>

<div style="text-align: center;">

[![Hexagonal Architecture](https://img.shields.io/badge/Architecture-Hexagonal-40B5A4.svg)](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))
[![Docker](https://img.shields.io/badge/Container-Docker-2496ED.svg)](https://www.docker.com/)

[![Apache 2.0 License](https://img.shields.io/badge/License-Apache%202.0-orange)](./LICENSE)
![GO 1.22](https://img.shields.io/badge/go-1.22-blue)
</div>

This project is a Golang-based application for notifications, utilizing a Hexagonal Architecture pattern for a clean separation of concerns and modular design.

## Project Structure

The application is organized around the main entity: `notification`. Each entity is encapsulated within three packages representing the layers of the Hexagonal Architecture:

- `infrastructure`: Contains code related to database operations, networking, and interfacing with external systems.
- `application`: Hosts the application logic, defining use cases and orchestrating between the infrastructure and domain layers.
- `domain`: Defines the business models and core business logic.

The following diagram illustrates the Hexagonal Architecture pattern:

- ![hexagonal.webp](docs/img/hexagonal.webp)

## Technologies Used

- **Programming Language:** GO 1.22
- **Build Tool:** Build

## How to Use

start container

```docker-compose
docker-compose up --build
```
