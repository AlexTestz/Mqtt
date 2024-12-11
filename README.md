

# Hello World  MQTT in GO  - Docker Image

This Docker image contains a simple go program using  that responds with "Hello, World!" when accessed via an HTTP GET request.

## Contents of the Image

The image includes the following files and configurations:

- **main.go***: The main go script that defines the  application and sets up the endpoint.
- **Dockerfile**: Configuration file that defines the Docker image, sets up the go environment, and specifies the instructions to run the program when the container starts.
- **Go**: Uses the base go image, version 3.11-slim.

## How to Use This Image

To run the program on your machine, make sure Docker is installed. Then, execute the following command in your terminal:


             docker run alexmpz/hola-mqtt

