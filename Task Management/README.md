
## Task Management API - README

Welcome to the Task Management API, a RESTful API built using the Go Gin framework. This project provides endpoints for managing tasks, including creating, retrieving, updating, and deleting tasks.


### Table of Contents

1. [Overview](#overview)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Running the API](#running-the-api)
5. [API Endpoints](#api-endpoints)
6. [License](#license)

### Overview

The Task Management API is designed to be simple and efficient, providing a way to manage tasks via a REST API. It uses the Gin framework for handling HTTP requests and responses. The application is a self-contained project and does not require a database; all data is stored in memory.

### Prerequisites

Before setting up the project, ensure you have the following installed on your machine:

1. **Go**: Version 1.16 or higher. Download it from [golang.org](https://golang.org/dl/).
2. **Git**: For version control.

### Installation

#### 1. Clone the Repository

First, clone the repository to your local machine:

```sh
git clone https://github.com/saleamlakw/A2SV_backend_track.git
cd A2SV_backend_track/Task Management
```

#### 2. Install Dependencies

Install the necessary Go packages:

```sh
go mod tidy
```

This command will download all dependencies specified in the `go.mod` file.

### Running the API

To start the API server, use the following command:

```sh
go run main.go
```

The server will start on `http://localhost:8080` by default. If you've set a different port using environment variables, it will use that port instead.

### API Endpoints

The following endpoints are available in the Task Management API:

1. **Get All Tasks**
   - **Endpoint**: `/tasks`
   - **Method**: `GET`
   - **Description**: Retrieves all tasks.
   - **Response**: JSON array of task objects.

2. **Create a New Task**
   - **Endpoint**: `/tasks`
   - **Method**: `POST`
   - **Description**: Creates a new task.
   - **Request Body**: JSON object containing task details.
   - **Response**: JSON object of the created task.

3. **Get Task by ID**
   - **Endpoint**: `/tasks/:id`
   - **Method**: `GET`
   - **Description**: Retrieves a task by its ID.
   - **Response**: JSON object of the task, or a 404 error if not found.

4. **Update Task**
   - **Endpoint**: `/tasks/:id`
   - **Method**: `PUT`
   - **Description**: Updates an existing task by ID.
   - **Request Body**: JSON object with updated task details.
   - **Response**: JSON object of the updated task, or a 404 error if not found.

5. **Delete Task**
   - **Endpoint**: `/tasks/:id`
   - **Method**: `DELETE`
   - **Description**: Deletes a task by its ID.
   - **Response**: JSON message confirming deletion, or a 404 error if not found.

### License

This project is licensed under the MIT License. For more details, refer to the [LICENSE](LICENSE) file in the repository.
