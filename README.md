# Chat App

## Description
A simple chat application built with Golang and WebSockets.

## Prerequisites
- Docker

## Installation and Running

1. Clone the repository:

    ```bash
    git clone git@github.com:pramudya3/chat-app.git
    cd chat-app
    ```

2. Build the Docker image:

    ```bash
    docker build -t chat-app .
    ```

3. Run the Docker container:

    ```bash
    docker run -p 8080:8080 chat-app
    ```

4. Access the application at `http://localhost:8080`.

## Using Docker Compose
You can also use Docker Compose to build and run the application:

1. Build and run the application:

    ```bash
    docker compose build && docker compose up
    ```

2. Access the application at `http://localhost:8080`.

## Trying Out the Chat Application
To try chatting with this project:

1. Open 2 new tabs in your web browser.
2. In each tab, type `localhost:8080`.
3. Insert your name in each tab.
4. Start typing messages and see them appear in both tabs.

## License
This project is licensed under the MIT License.
