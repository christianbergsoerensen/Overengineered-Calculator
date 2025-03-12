# Overengineered-Calculator

This is a simple RESTful API that performs arithmetic operations and stores the calculation history. It supports addition, subtraction, multiplication, and division operations.

## Features:
- Perform basic arithmetic operations: addition, subtraction, multiplication, and division.
- View calculation history.

## Operations Supported:
- **POST /calculate**: Perform an arithmetic operation.
    - Request body (JSON format):
        ```json
        {
          "operation": "add",
          "a": 10,
          "b": 5
        }
        ```
    - Valid operations: `add`, `subtract`, `multiply`, `divide`.
    - Response: The result of the operation.

- **GET /history**: Retrieve the history of past calculations.
    - Response: List of all calculations with id, operation, operands, result and timestamp.

## Running the API Locally

### Steps:
1. Clone the repository:
   ```bash
   git clone github.com/christianbergsoerensen/Overengineered-Calculator
   cd Overengineered-Calculator

2. Install dependencies:
   ```bash
   go mod tidy

3. Build and run the application:
    ```bash
   go run cmd/main.go