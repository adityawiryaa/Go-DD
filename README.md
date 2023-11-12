# Go DD

Welcome to Go DD, a state-of-the-art application leveraging Domain-Driven Design (DDD), Table-Driven Testing, and Test-Driven Development (TDD). Offers a robust and scalable platform, ideal for businesses looking to adopt modern software practices.

## Features

- **Domain-Driven Design**: Focused around domain logic to meet complex business requirements with clarity.
- **Table-Driven Tests**: Enhance test quality and coverage, making them more readable and extendable.
- **Test-Driven Development**: Reliability and bug resistance built-in with TDD methodologies.

## Getting Started

These instructions will get your copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Ensure you have the following installed before running the application:
- Go (Version 1.21.1)
- Relevant environment setup and dependencies

### Installation
Clone the repository:

```bash
git clone https://github.com/adityawiryaa/go-driven-design.git
cd go-driven-design
go mod tidy
```
### Build the binary
Just run `make build` it build the binary named `go-driven-design`
### Running the Service
Start the service using the following command:
```bash
make start
```
This command starts the service on localhost:3000.
### Running the Tests
Run the tests with:
```bash
make test
```
These tests, designed with TDD principles, ensure the application's reliability and efficiency.
## API Usage
Interact with the API by sending a POST request to `localhost:3000` with the following JSON payload:

```json
{
    "first_name": "Jane",
    "last_name": "Doe"
}
```
The API will respond with a personalized greeting, like "Hello, Jane Doe".