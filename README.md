# 3-Pass Compiler

Microservice Architecture Development of a Three-Pass Compiler

## Phase 1: Planning and Design

##### Define Language Subset

    > Clearly document the subset of C++ language features you’ll support. This should include:
        - Data types
            - int, float, char, bool.
        - Operators
            - Arithmetic (+, -, *, /)
            - Logical (&&, ||, !)
            - Comparison (==, !=, >, <, <=, >=)
        - Control structures
            - if statements, for loops, while loops.
        - Functions
            - Simple function definitions and calls, basic return values.
        - Input/Output
            - cin, cout.

##### System Architecture Design

    > Finalize the architecture diagram showing each microservice’s role (Lexical Analysis, Parsing, Code Generation) and how data flows between them.
    > Document each microservice's API specifications, including endpoints, input/output formats, and data flow.

##### Error Handling Strategy

    > Define an error-handling approach for each phase:
        - Lexical Analysis Errors (Tokenizer)
            - Illegal tokens, unrecognized symbols, improper syntax.
        - Parsing Errors (AST)
            - Unmatched braces, missing semi-colons, undefined variable references.
        - Code Generation Errors (Assembly Generator)
            - Unsupported constructs, undefined operations, memory issues.
    > Ensure error messages are clear and standardized across services.

##### Development Tools and Language Choice

    > Confirm development language for each microservice API
        - Go.
    > Confirm Docker setup is appropriate for each microservice.
        - Yes.
    > Choose Infrastructure as Code (IaC) tool(s) like Docker Compose, Terraform, or Ansible for deployment.
        - Docker Compose.

##### Infrastructure as Code (IaC) Planning

    > Plan Docker networking to ensure each microservice can communicate using TCP sockets.
    > Draft Docker Compose file or Terraform configuration to manage the orchestration.

## Phase 2: Microservice Development

##### Lexical Analysis Microservice

    > Develop a tokenizer that breaks down source code into tokens based on the language subset.
    > Create an API endpoint that accepts source code input and outputs tokens as a JSON object.
    > Write a Dockerfile and containerize the Lexical Analysis microservice.

##### Parsing Microservice

    > Build a parser that constructs an Abstract Syntax Tree (AST) from tokens, following C++ grammar rules.
    > Create an API endpoint that accepts tokens as input and returns the AST as JSON.
    > Write a Dockerfile and containerize the Parsing microservice.

##### Code Generation Microservice

    > Develop code generation logic that converts the AST into machine code or assembly-like instructions.
    > Create an API endpoint that accepts an AST and returns the generated machine code.
    > Write a Dockerfile and containerize the Code Generation microservice.

##### Networking and Communication Testing

    > Set up a Docker network for inter-service communication using TCP.
    > Test socket-based communication between each service (token transfer, AST transfer).

## Phase 3: Container Communication and Networking

##### Network Configuration

    > Finalize and test the Docker network setup, ensuring all microservices can communicate over TCP.

##### Inter-Service API Communication

    > Implement inter-service data transfer using network sockets (e.g., sending tokens from Lexical Analysis to Parsing).
    > Verify JSON data format consistency between microservices and troubleshoot any serialization issues.

##### Integration Testing

    > Test end-to-end compilation from source code input to final machine code output across all three services.
    > Validate error handling between services, checking that each phase reports errors back to the user correctly.

## Phase 4: Orchestration and Deployment Automation

##### Infrastructure as Code (IaC) Setup

    > Use Docker Compose, Terraform, or Ansible to automate container deployment and teardown.
    > Configure volumes, environment variables, and port mappings in Docker Compose or Terraform.

##### Platform Testing

    > Test deployment across different operating systems (Windows, macOS, Linux) with Docker installed to ensure platform independence.

##### Documentation of Deployment Process

    > Document deployment and teardown instructions using IaC tools, making it easy for users to set up the compiler on any Docker-compatible system.

## Phase 5: User Interface and Testing

##### User Interface Setup

    > Develop a basic interface (CLI, web, or API client) for users to submit source code, trigger compilation, and download the final machine code output.

##### Functional and Performance Testing

    > Run tests with varying input code sizes (50-500 lines) to ensure the compilation process completes within the five-second requirement.
    > Check error messages for clarity, consistency, and usability.

##### Documentation

    > Write documentation on using the compiler, setting up the environment, API usage, and troubleshooting.
    > Finalize API documentation for each service, describing endpoints, data formats, and error messages.
