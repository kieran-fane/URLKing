# URLKing

URLKing is a lightweight URL shortener application built to conquer imposter syndrome after a challenging interview experience. It features a simple, intuitive frontend and a flexible backend supporting both Dockerized PostgreSQL and SQLite, enabling quick setup and scalability.

---

## Features

- **Shorten URLs**: Instantly generate compact, shareable links.  
- **Redirect**: Seamlessly redirect from short URLs to original destinations.  
- **Flexible Backend**:  
  - **Dockerized PostgreSQL** for production-like environments.  
  - **SQLite** for quick local development without external dependencies.  
- **Minimal Frontend**: Clean UI to submit and manage URLs.  

---

## Technologies

- **Frontend**: React, Vite, npm  
- **Backend**: Go  
  - PostgreSQL (via Docker)  
  - SQLite  
- **Containerization**: Docker  

---

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/)  
- [Go](https://golang.org/) (v1.16+)  
- [Node.js](https://nodejs.org/) & npm  

### Backend

#### Dockerized PostgreSQL

1. Open a terminal and navigate to the Docker backend folder:  
   ```bash
   cd backend/docker
   ./run.sh
   ```
2. The backend API will be running on `http://localhost:8080`.

#### SQLite (Local Development)

1. Open a terminal and navigate to the SQLite backend folder:
  ```bash
    cd backend/sqlite
    go run server.go
  ```
2. The backend API will be running on `http://localhost:8080`.

## Frontend

1. Open a terminal and navigate to the frontend folder:  
   ```bash
   cd frontend
   npm install
   npm run dev
   ```
2. The app will open automatically at `http://localhost:3000`.

### Usage
- Access the frontend in your browser.

- Enter a URL in the input field and click Shorten.

- Copy or click the generated short link to be redirected to the original URL.

*Built with ❤️ to remind myself that growth comes from every experience — even the tough ones.*
