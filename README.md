# Social Network Platform

Welcome to the Social Network Platform project! This project aims to create a fully functional social networking website where users can create accounts, log in, view other people's posts, add friends, and chat with friends. This platform will be developed with a focus on both frontend and backend technologies, ensuring a seamless user experience and robust backend support.

## Features

- **User Authentication**: Users can create an account and log in.
- **User Posts**: Users can create, view, and interact with posts.
- **Friendship Management**: Users can add and manage friends.
- **Messaging System**: Users can chat with their friends in real-time.
- **Responsive Design**: The website will be responsive and user-friendly across different devices.

## Technologies Used

### Frontend

- **HTML**
- **CSS**
- **JavaScript**

### Backend

- **Golang**

### Database

- **MariaDB**

### Containerization

- **Docker** (for containerizing the application upon completion)

## Project Structure

The project is divided into several key components:

1. **Frontend**: This includes all the client-side code responsible for rendering the user interface and handling user interactions. The frontend is built using HTML, CSS, and JavaScript.

2. **Backend**: The server-side logic is handled by Golang, which includes user authentication, post management, friend management, and messaging functionality.

3. **Database**: All data is stored in a MariaDB database, ensuring reliable and efficient data management.

4. **Containerization**: Once the project is complete, it will be containerized using Docker to ensure easy deployment and scalability.

## Getting Started

### Prerequisites

- **Golang**: Ensure Golang is installed on your machine.
- **Node.js**: Required for frontend dependencies.
- **MariaDB**: Install MariaDB for database management.
- **Docker**: For containerization.

### Installation

1. **Clone the Repository**

   ```sh
   git clone git@github.com:YourUsername/All-Chat.git
   cd All-Chat
    ``
2. **Set Up the Backend**

- Navigate to the backend directory:

  ```sh
  cd backend
  ```
- Install Golang dependencies:
    ```sh
    go mod tidy
    ```
- Set up the MariaDB database and update the connection settings in the backend configuration.


3 **Set Up the Frontend**
- Navigate to the frontend directory:
    ```sh
    Copy code
    cd frontend
    ```
- Install frontend dependencies:
    ```sh
    Copy code
    npm install
    ```

4 **Run the Application**
- Start the backend server:
    ```sh
    Copy code
    go run main.go
    ```
- Serve the frontend:
    ```sh
    Copy code
    npm start
    ```

### Containerization
Once the application is ready for deployment, it will be containerized using Docker. Instructions for building and running Docker containers will be provided at that time.