# Go WebSocket Chat

This project is a simple WebSocket chat application built with Go. It allows users to connect via WebSockets, send and receive messages in real-time.

## Technologies Used

- **Go**: The programming language used to build the server.
- **WebSocket**: The protocol used for real-time communication between the client and server.
- **Docker**: Used to containerize the application for easy deployment.
- **Testing**: The project includes tests to ensure functionality.

## Getting Started

To run this application, ensure you have Docker and Docker Compose installed on your machine.

### Clone the Repository

```bash
git clone https://github.com/yebrai/Go-ws-chat
cd <repository-directory>
```
### Build and Run the Application
To build and start the application, use the following command:

```bash
make build
```
This command will build the Docker image and start the container.

### Running Tests
To run the tests for the application, execute:

```bash
make test
```
### Accessing the Application
Once the application is running, you can access it at:

```
http://localhost:3000
```

### Clean Up
To remove the Docker image after you're done, you can run:

```
make clean
```