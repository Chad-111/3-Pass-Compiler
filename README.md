# Three-Pass Compiler with Microservice Architecture

## Overview

A three-pass compiler for a C++ language subset, implemented using microservices for Lexical Analysis, Parsing, and Code Generation.

## Supported Language Features

- **Data Types:** `int`, `float`, `char`, `bool`
- **Operators:** Arithmetic (`+`, `-`, `*`, `/`), Logical (`&&`, `||`, `!`), Comparison (`==`, `!=`, `>`, `<`, `<=`, `>=`)
- **Control Structures:** `if`, `for`, `while`
- **Functions:** Basic function definitions, calls, and return values
- **I/O:** `cin`, `cout`

## System Design

### Architecture
- **Microservices:** 
  - **Lexical Analysis:** Tokenizer that breaks source code into tokens.
  - **Parsing:** Constructs an Abstract Syntax Tree (AST) from tokens.
  - **Code Generation:** Converts the AST into machine code or assembly-like instructions.
- **Communication:** TCP-based communication between services using JSON for data exchange.

### Error Handling
- **Lexical Analysis:** Illegal tokens, unrecognized symbols
- **Parsing:** Unmatched braces, missing semicolons, undefined variables
- **Code Generation:** Unsupported constructs, undefined operations, memory issues
- **Standardized Messages:** Consistent and clear error feedback across services

## Development Environment

- **Language:** Go
- **Containerization:** Docker with Docker Compose
- **Networking:** TCP sockets for inter-service communication

## Phases

### Phase 1: Planning and Design
- Define language subset
- Design architecture diagram and API specifications
- Establish error-handling strategy and standard messages

### Phase 2: Microservice Development
- **Lexical Analysis:** Develop tokenizer, create API endpoint, and containerize
- **Parsing:** Construct AST, create API endpoint, and containerize
- **Code Generation:** Generate machine code, create API endpoint, and containerize

### Phase 3: Networking and Communication
- Finalize Docker network and test TCP-based inter-service communication
- Ensure consistent JSON format for data exchange

### Phase 4: Orchestration and Deployment
- Use Docker Compose for automated deployment and setup across platforms
- Document IaC setup, including volumes, environment variables, and port mappings

### Phase 5: User Interface and Testing
- **Interface:** CLI or web API client for submitting source code and retrieving compiled output
- **Testing:** Validate performance (compilation within five seconds) and consistency of error messages
- **Documentation:** Provide usage instructions, API details, and troubleshooting guidance

## Getting Started

1. **Setup:** Clone the repository, install Docker, and run `docker-compose up` to start all services.
2. **Usage:** Submit code via the interface, check for errors, and retrieve final machine code output.
3. **Documentation:** Refer to detailed API docs for each service and deployment instructions.

## Contact and Support

For setup issues or questions, refer to the documentation or submit issues via GitHub.
