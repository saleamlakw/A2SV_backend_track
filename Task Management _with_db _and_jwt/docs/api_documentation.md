# Project: Task Manager
The Task Manager API allows users to efficiently manage tasks, providing endpoints for creating, retrieving, updating, and deleting tasks. It is designed to facilitate task tracking and management for individuals and teams.

## End-point: list tasks
This endpoint retrieves a list of tasks from the server.

### Request

- Method: GET
    
- URL: `localhost:8080/tasks`
    

### Response

- Status: 200
    
- Content-Type: application/json
    
- \[ { "id": "", "title": "", "description": "", "due_date": "", "status": "" } \]
    

The response contains an array of objects, where each object represents a task and includes the following parameters:

- `id` (string): The unique identifier for the task.
    
- `title` (string): The title or name of the task.
    
- `description` (string): A description of the task.
    
- `due_date` (string): The due date for the task.
    
- `status` (string): The status of the task.
### Method: GET
>```
>localhost:8080/tasks
>```
### Response: 200
```json
[
    {
        "id": "1",
        "title": "Task 1",
        "description": "First task",
        "due_date": "2024-08-01T09:51:12.1768146+03:00",
        "status": "Pending"
    },
    {
        "id": "2",
        "title": "Task 2",
        "description": "Second task",
        "due_date": "2024-08-02T09:51:12.1768146+03:00",
        "status": "In Progress"
    },
    {
        "id": "3",
        "title": "Task 3",
        "description": "Third task",
        "due_date": "2024-08-03T09:51:12.1778746+03:00",
        "status": "Completed"
    }
]
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: get task by id
### Request

This endpoint makes an HTTP GET request to retrieve data.

### Request Body

This request does not require a request body.

### Response

The response for this request is a JSON object with the following schema:

- id (string): The unique identifier for the item.
    
- title (string): The title of the item.
    
- description (string): The description of the item.
    
- due_date (string): The due date for the item.
    
- status (string): The status of the item.
### Method: GET
>```
>localhost:8080/tasks/1
>```
### Response: 200
```json
{
    "id": "1",
    "title": "Task 1",
    "description": "First task",
    "due_date": "2024-08-01T09:51:12.1768146+03:00",
    "status": "Pending"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: post task
This endpoint allows you to create a new task by sending an HTTP POST request to localhost:8080/tasks. The request should include a JSON payload in the raw request body with the following parameters:

- id (string): The unique identifier for the task.
    
- title (string): The title of the task.
    
- description (string): A description of the task.
    
- due_date (string): The due date for the task.
    
- status (string): The status of the task.
    

### Request Body

``` json
{
  "id": "",
  "title": "",
  "description": "",
  "due_date": "",
  "status": ""
}

 ```

#### Example Response

Upon successful creation of the task, the endpoint will respond with a JSON object representing the newly created task, including the following parameters:

- id (string): The unique identifier for the task.
    
- title (string): The title of the task.
    
- description (string): A description of the task.
    
- due_date (string): The due date for the task.
    
- status (string): The status of the task.
    

``` json
{
  "id": "",
  "title": "",
  "description": "",
  "due_date": "",
  "status": ""
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
    "id": "5",
    "title": "Task 5",
    "description": "First task",
    "due_date": "2024-08-01T09:51:12.1768146+03:00",
    "status": "Pending"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: delete task by id
### Delete Task

This endpoint is used to delete a specific task identified by its ID.

#### Request Body

This request does not require a request body.

#### Response

- `message` (string): A message indicating the result of the deletion operation.
### Method: DELETE
>```
>localhost:8080/tasks/1
>```
### Body (**raw**)

```json

```

### Response: 200
```json
{
    "message": "task removed"
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

The response for this request is a JSON object with a `message` key, which provides information about the success or failure of the update operation.

#### JSON Schema for Response

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
### Method: PUT
>```
>localhost:8080/tasks/2
>```
### Body (**raw**)

```json
{
    "id": "2",
    "title": "Task tata",
    "description": "First task sik",
    "due_date": "2024-08-01T09:51:12.1768146+03:00",
    "status": "Pending"
}
```

### Response: 200
```json
{
    "message": "task updated sucessfully"
}
```
⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
