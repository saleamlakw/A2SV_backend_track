# Project: Task Manager
The Task Manager API allows users to efficiently manage tasks, providing endpoints for creating, retrieving, updating, and deleting tasks. It is designed to facilitate task tracking and management for individuals and teams.

## End-point: list tasks
### GET /tasks

This endpoint retrieves a list of tasks.

#### Request

This request does not require any request body.

#### Response

The response will be in JSON format with the following schema:

``` json
[
  {
    "_id": "",
    "title": "",
    "description": "",
    "due_date": "",
    "status": "",
    "user_id": ""
  }
]

 ```

The response will contain an array of task objects, where each task object will have properties like _id, title, description, due_date, status, and user_id.
### Method: GET
>```
>localhost:8080/tasks
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxMjYwOTksInJvbGUiOiJhZG1pbiIsInVzZXJpZCI6IjY2YjI4YWY4MDExNGU1OGMyMDk3NDNkOSIsInVzZXJuYW1lIjoibmFuYSJ9.UtOa2P1wzoeNHnyZU7UPqFiBGacQP4WyXvB-46h_Uy4|


### Response: 200
```json
[
    {
        "_id": "66b35b27e60f5fc0ec9eb6b8",
        "title": "Task 10",
        "description": "First task",
        "due_date": "2024-08-01T06:51:12.176Z",
        "status": "Pending",
        "user_id": "66b28c1c7f13c3b67c9b1ad9"
    }
]
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: get task by id
### Request

This endpoint makes an HTTP GET request to retrieve data for a specific task.

### Request Body

This request does not require a request body.

### Response

The response for this request is a JSON object with the following schema:

- _id (string): The unique identifier for the task.
    
- title (string): The title of the task.
    
- description (string): The description of the task.
    
- due_date (string): The due date for the task.
    
- status (string): The status of the task.
    
- user_id (string): The unique identifier of the user associated with the task.
### Method: GET
>```
>localhost:8080/tasks/66b28c1c7f13c3b67c9b1ad9
>```
### Headers

|Content-Type|Value|
|---|---|
|Cache-Control|no-cache|


### Headers

|Content-Type|Value|
|---|---|
|Postman-Token|<calculated when request is sent>|


### Headers

|Content-Type|Value|
|---|---|
|Host|<calculated when request is sent>|


### Headers

|Content-Type|Value|
|---|---|
|User-Agent|PostmanRuntime/7.39.1|


### Headers

|Content-Type|Value|
|---|---|
|Accept|*/*|


### Headers

|Content-Type|Value|
|---|---|
|Accept-Encoding|gzip, deflate, br|


### Headers

|Content-Type|Value|
|---|---|
|Connection|keep-alive|


### Headers

|Content-Type|Value|
|---|---|
|Authorization|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxMTY5MjQsInJvbGUiOiJ1c2VyIiwidXNlcmlkIjoiNjZiMjllOWQ5ZGJjYzhjMTc5MTYxNWFmIiwidXNlcm5hbWUiOiJsYWtldyJ9.Ww12qouFMEwFjEIqVFJOUsLKXZzrKwDTm7tDJeGxlEA|


### Headers

|Content-Type|Value|
|---|---|
|||


### Response: 200
```json
{
    "_id": "66b35b27e60f5fc0ec9eb6b8",
    "title": "Task 10",
    "description": "First task",
    "due_date": "2024-08-01T06:51:12.176Z",
    "status": "Pending",
    "user_id": "66b28c1c7f13c3b67c9b1ad9"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: post task
This endpoint allows you to create a new task by sending an HTTP POST request to localhost:8080/tasks. The request should include a JSON payload in the raw request body with the following parameters:

- title (string): The title of the task.
    
- description (string): A description of the task.
    
- due_date (string): The due date for the task.
    
- status (string): The status of the task.
    
- user_id (string): The unique identifier of the user associated with the task.
    

### Request Body

``` json
{
  "title": "",
  "description": "",
  "due_date": "",
  "status": "",
  "user_id": ""
}

 ```

#### Example Response

Upon successful creation of the task, the endpoint will respond with a JSON object representing the newly created task, including the following parameters:

- _id (string): The unique identifier for the task.
    
- title (string): The title of the task.
    
- description (string): A description of the task.
    
- due_date (string): The due date for the task.
    
- status (string): The status of the task.
    
- user_id (string): The unique identifier of the user associated with the task.
    

``` json
{
  "_id": "",
  "title": "",
  "description": "",
  "due_date": "",
  "status": "",
  "user_id": ""
}

 ```
### Method: POST
>```
>localhost:8080/tasks
>```
### Body (**raw**)

```json
{
    "id": "5",
    "title": "Task 5",
    "description": "First task",
    "due_date": "2024-08-01T09:51:12.1768146+03:00",
    "status": "Pending"
}
```

### Response: 201
```json
{
    "_id": "66b382bb18bac5570b011dc7",
    "title": "dangling",
    "description": "First task",
    "due_date": "2024-08-01T09:51:12.1768146+03:00",
    "status": "done",
    "user_id": "66b29e9d9dbcc8c1791615af"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: delete task by id
### Delete Task

This endpoint is used to delete a specific task identified by its ID.

#### Request Body

This request does not require a request body.

#### Response Body

- `message` (string): A message indicating the result of the deletion operation.
### Method: DELETE
>```
>localhost:8080/tasks/66b35a4c09402d61d1d504a6
>```
### Headers

|Content-Type|Value|
|---|---|
|Cache-Control|no-cache|


### Headers

|Content-Type|Value|
|---|---|
|Postman-Token|<calculated when request is sent>|


### Headers

|Content-Type|Value|
|---|---|
|Host|<calculated when request is sent>|


### Headers

|Content-Type|Value|
|---|---|
|User-Agent|PostmanRuntime/7.39.1|


### Headers

|Content-Type|Value|
|---|---|
|Accept|*/*|


### Headers

|Content-Type|Value|
|---|---|
|Accept-Encoding|gzip, deflate, br|


### Headers

|Content-Type|Value|
|---|---|
|Connection|keep-alive|


### Headers

|Content-Type|Value|
|---|---|
|Authorization|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxMTcyMjcsInJvbGUiOiJhZG1pbiIsInVzZXJpZCI6IjY2YjI4YWY4MDExNGU1OGMyMDk3NDNkOSIsInVzZXJuYW1lIjoibmFuYSJ9.9mRkaZmFrnijJu4PAeiSBdEapLxY6l3OZ8pAIdRlL00|


### Body (**raw**)

```json

```

### Response: 200
```json
{
    "message": "task deleted"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: update task by id
### Update Task

This endpoint allows the client to update a specific task by sending a PUT request to the specified URL.

#### Request Body

- `id` (string) - The unique identifier of the task.
    
- `title` (string) - The title of the task.
    
- `description` (string) - The description of the task.
    
- `due_date` (string) - The due date of the task.
    
- `status` (string) - The status of the task.
    

#### Response

The response for this request is a JSON object with the following properties:

- `_id` (string) - The unique identifier of the task.
    
- `title` (string) - The title of the task.
    
- `description` (string) - The description of the task.
    
- `due_date` (string) - The due date of the task.
    
- `status` (string) - The status of the task.
    
- `user_id` (string) - The unique identifier of the user associated with the task.
    

#### JSON Schema for Response

``` json
{
    "type": "object",
    "properties": {
        "_id": {
            "type": "string"
        },
        "title": {
            "type": "string"
        },
        "description": {
            "type": "string"
        },
        "due_date": {
            "type": "string"
        },
        "status": {
            "type": "string"
        },
        "user_id": {
            "type": "string"
        }
    }
}

 ```
### Method: PUT
>```
>localhost:8080/tasks/66b35a4c09402d61d1d504a6
>```
### Headers

|Content-Type|Value|
|---|---|
|Cache-Control|no-cache|


### Headers

|Content-Type|Value|
|---|---|
|Postman-Token|<calculated when request is sent>|


### Headers

|Content-Type|Value|
|---|---|
|Content-Type|text/plain|


### Headers

|Content-Type|Value|
|---|---|
|Content-Length|<calculated when request is sent>|


### Headers

|Content-Type|Value|
|---|---|
|Host|<calculated when request is sent>|


### Headers

|Content-Type|Value|
|---|---|
|User-Agent|PostmanRuntime/7.39.1|


### Headers

|Content-Type|Value|
|---|---|
|Accept|*/*|


### Headers

|Content-Type|Value|
|---|---|
|Accept-Encoding|gzip, deflate, br|


### Headers

|Content-Type|Value|
|---|---|
|Connection|keep-alive|


### Headers

|Content-Type|Value|
|---|---|
|Authorization|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxMTcyMjcsInJvbGUiOiJhZG1pbiIsInVzZXJpZCI6IjY2YjI4YWY4MDExNGU1OGMyMDk3NDNkOSIsInVzZXJuYW1lIjoibmFuYSJ9.9mRkaZmFrnijJu4PAeiSBdEapLxY6l3OZ8pAIdRlL00|


### Body (**raw**)

```json
{
    "title": "laporte",
    "description": "First task sik",
    "due_date": "2024-08-01T09:51:12.1768146+03:00",
    "status": "done"
}
```

### Response: 200
```json
{
    "_id": "66b382bb18bac5570b011dc7",
    "title": "task 6",
    "description": "it was a bit challenging",
    "due_date": "2024-08-01T06:51:12.176Z",
    "status": "done",
    "user_id": "66b29e9d9dbcc8c1791615af"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: createUser
This endpoint allows users to sign up by sending a POST request to localhost:8080/users/signup. The request should include a JSON payload in the raw request body with "username" and "password" fields.

### Request Body

- `username` (string, required): The username of the user.
    
- `password` (string, required): The password for the user account.
    

### Response

The response to this request is in the form of a JSON schema with the following properties:

- `_id` (string): The unique identifier for the user.
    
- `username` (string): The username of the user.
    
- `password` (string): The password of the user.
    
- `role` (string): The role assigned to the user.
    

#### Example Response

``` json
{
  "_id": "",
  "username": "",
  "password": "",
  "role": ""
}

 ```
### Method: POST
>```
>localhost:8080/users/signup
>```
### Body (**raw**)

```json
{
"username": "sale",
"password": "nadew"
}
```

### Response: 201
```json
{
    "_id": "66b38502b71a7e4385ee39c6",
    "username": "sala",
    "password": "$2a$10$Ot7EqD38Cx8R8XAh6oQsDem034bHYxYxapWt1/2elaZ6Z.atyGJEG",
    "role": "user"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: login
### POST /users/login

This endpoint is used to authenticate users and log them into the system.

#### Request Body

- `username` (string, required): The username of the user.
    
- `password` (string, required): The password of the user.
    

#### Response

The response is in JSON format with the following schema:

``` json
{
    "type": "object",
    "properties": {
        // Add properties here
    }
}

 ```
### Method: POST
>```
>localhost:8080/users/login
>```
### Body (**raw**)

```json
{
"username": "nana",
"password": "Firsttask"
}
```

### Response: 200
```json
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxMjYwOTksInJvbGUiOiJhZG1pbiIsInVzZXJpZCI6IjY2YjI4YWY4MDExNGU1OGMyMDk3NDNkOSIsInVzZXJuYW1lIjoibmFuYSJ9.UtOa2P1wzoeNHnyZU7UPqFiBGacQP4WyXvB-46h_Uy4"
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: promote
The endpoint sends an HTTP POST request to localhost:8080/promote/66b38502b71a7e4385ee39c6 in order to promote a specific item. The response returned has a status code of 200 and a content type of application/json. The response body contains a JSON object with a key "message" which may have an empty string value. Below is the JSON schema for the response:

``` json
{
  "type": "object",
  "properties": {
    "message": {
      "type": "string"
    }
  }
}

 ```
### Method: POST
>```
>localhost:8080/promote/66b28c1c7f13c3b67c9b1ad9
>```
### Headers

|Content-Type|Value|
|---|---|
|Cache-Control|no-cache|


### Headers

|Content-Type|Value|
|---|---|
|Postman-Token|<calculated when request is sent>|


### Headers

|Content-Type|Value|
|---|---|
|Host|<calculated when request is sent>|


### Headers

|Content-Type|Value|
|---|---|
|User-Agent|PostmanRuntime/7.39.1|


### Headers

|Content-Type|Value|
|---|---|
|Accept|*/*|


### Headers

|Content-Type|Value|
|---|---|
|Accept-Encoding|gzip, deflate, br|


### Headers

|Content-Type|Value|
|---|---|
|Connection|keep-alive|


### Headers

|Content-Type|Value|
|---|---|
|Authorization|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxMTcyMjcsInJvbGUiOiJhZG1pbiIsInVzZXJpZCI6IjY2YjI4YWY4MDExNGU1OGMyMDk3NDNkOSIsInVzZXJuYW1lIjoibmFuYSJ9.9mRkaZmFrnijJu4PAeiSBdEapLxY6l3OZ8pAIdRlL00|


### Response: 200
```json
{
    "message": "role updated sucessfully"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: getUserById
This endpoint retrieves user information based on the provided user ID. The response returns a JSON object with the user's details including the user ID, username, password, and role.

``` json
{
  "type": "object",
  "properties": {
    "_id": {
      "type": "string"
    },
    "username": {
      "type": "string"
    },
    "password": {
      "type": "string"
    },
    "role": {
      "type": "string"
    }
  }
}

 ```
### Method: GET
>```
>
>```
### Response: 200
```json
{
    "_id": "66b3591209402d61d1d504a2",
    "username": "abebe",
    "password": "$2a$10$1O7VmAe9G1lylD/m7Y.S1eZj2QIqpLnd302/mu8XoPTPtUwXXBvDC",
    "role": "user"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: getUsers
This endpoint makes an HTTP GET request to localhost:8080/users to retrieve a list of users. The response will be in JSON format with an array of user objects, each containing the user's ID, username, password, and role.

### Request Body

This request does not require a request body.

### Response Body

- `_id` (string): The unique identifier for the user.
    
- `username` (string): The username of the user.
    
- `password` (string): The password of the user.
    
- `role` (string): The role of the user within the system.
    

#### Example

``` json
[
    {
        "_id": "",
        "username": "",
        "password": "",
        "role": ""
    }
]

 ```
### Method: GET
>```
>localhost:8080/users
>```
### Headers

|Content-Type|Value|
|---|---|
|Cache-Control|no-cache|


### Headers

|Content-Type|Value|
|---|---|
|Postman-Token|<calculated when request is sent>|


### Headers

|Content-Type|Value|
|---|---|
|Host|<calculated when request is sent>|


### Headers

|Content-Type|Value|
|---|---|
|User-Agent|PostmanRuntime/7.39.1|


### Headers

|Content-Type|Value|
|---|---|
|Accept|*/*|


### Headers

|Content-Type|Value|
|---|---|
|Accept-Encoding|gzip, deflate, br|


### Headers

|Content-Type|Value|
|---|---|
|Connection|keep-alive|


### Headers

|Content-Type|Value|
|---|---|
|Authorization|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxMjQ0MDEsInJvbGUiOiJhZG1pbiIsInVzZXJpZCI6IjY2YjI4YzFjN2YxM2MzYjY3YzliMWFkOSIsInVzZXJuYW1lIjoibmFuYXRpIn0.dabowufWIJhAiMwVjBwoaTC-JUIEMcUQVFrfz50Xpm4|


### Response: 200
```json
[
    {
        "_id": "66b28a0c0aae3ebc8f85385a",
        "username": "sale",
        "password": "$2a$10$duyACF/yCIhPZslGUXko6OkXetKoZ7t9E.ZoyOkEX6BdBIb/nwhZK",
        "role": "admin"
    },
    {
        "_id": "66b28af80114e58c209743d9",
        "username": "nana",
        "password": "$2a$10$u3C3q/5xELBCz8YL2QSGTOuK6UdXWebAClvk.QsCgnS61wbH/J2uO",
        "role": "admin"
    },
    {
        "_id": "66b28c1c7f13c3b67c9b1ad9",
        "username": "nanati",
        "password": "$2a$10$Ba2LFhoBZE0ZJ04UMZhgEuOJS9DpAa4VxAxIAQ94E2akOUSRV3jda",
        "role": "admin"
    },
    {
        "_id": "66b29e9d9dbcc8c1791615af",
        "username": "lakew",
        "password": "$2a$10$YrKydE9Xye15NI2xCg1iaOefaROhNZM5wNPaa9HoLSBSJPf94NWPm",
        "role": "user"
    },
    {
        "_id": "66b3591209402d61d1d504a2",
        "username": "abebe",
        "password": "$2a$10$1O7VmAe9G1lylD/m7Y.S1eZj2QIqpLnd302/mu8XoPTPtUwXXBvDC",
        "role": "user"
    },
    {
        "_id": "66b38502b71a7e4385ee39c6",
        "username": "sala",
        "password": "$2a$10$Ot7EqD38Cx8R8XAh6oQsDem034bHYxYxapWt1/2elaZ6Z.atyGJEG",
        "role": "admin"
    }
]
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
_________________________________________________
