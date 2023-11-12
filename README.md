# TODOer

### Описание

TODOer - сайт для создания задач.
Пользователи могут создавать, читать, изменять и удалять задачи других
пользователей, добавлять рецепты в избранное и список покупок, подписываться
на других пользователей

### Технологии, использованые для реализации бэкенда

- Postgres 15
- go 1.20
-
    - [echo v4](https://echo.labstack.com/) (router)
-
    - [jackc/pgx](https://github.com/jackc/pgx) (db driver and connector)
-
    - [uber zap](https://github.com/uber-go/zap) (logger)
-
    - [uber fx](https://github.com/uber-go/fx) (DI container)
-
    - [easyjson](https://github.com/mailru/easyjson) (json serializing and deserializing) 
- Docker
- Docker Compose
- Tern (migrations)

### Запуск проекта в dev-режиме

- склонируйте репозиторий и перейдите в него

```
git clone https://github.com/vlad-marlo/todoer.git && cd todoer
```

- запустить docker-compose ```docker-compose up -d```

После этого сервер будет работать по адресу localhost:8080

### Документация API
# TODOer API
This is a todoer backend server.

## Version: 1.0

**Contact information:**  
API Support

### /tasks

#### POST
##### Summary:

Create task

##### Description:

Adds a new task to storage.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| request | body | Task object | Yes | [CreateTaskRequest](#CreateTaskRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [TaskDTO](#TaskDTO) |
| 400 | Bad Request | [ErrorMessage](#ErrorMessage) |
| 500 | Internal Server Error | [ErrorMessage](#ErrorMessage) |

#### GET
##### Summary:

Get tasks

##### Description:

Return a tasks with provided pagination and sort options.
If there is no tasks with provided options then empty array will be returned.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| offset | query | pagination offset of returned objects | No | integer |
| limit | query | pagination limit of returned objects | No | integer |
| task | query | Search parameter. Search tasks which contains provided string in it. | No | string |
| status | query | Tasks filtrated with statuses in this param. Statuses must be separated with commas. | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [GetTasksResponse](#GetTasksResponse) |

### /tasks/{task_id}

#### GET
##### Summary:

Get task

##### Description:

Return a task object by ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| task_id | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Returns a task by id | [TaskDTO](#TaskDTO) |
| 404 | There is no task with provided id | [ErrorMessage](#ErrorMessage) |
| 500 | Internal Server Error | [ErrorMessage](#ErrorMessage) |

#### PATCH
##### Summary:

Update task

##### Description:

Delete a task object by ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| task_id | path |  | Yes | string |
| task | body |  | Yes | [UpdateTaskRequest](#UpdateTaskRequest) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Task was updated | [TaskDTO](#TaskDTO) |
| 404 | Task does not exists |  |
| 500 | Internal Server Error | [ErrorMessage](#ErrorMessage) |

#### DELETE
##### Summary:

Delete task

##### Description:

Delete a task object by ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| task_id | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | Task was deleted |  |
| 500 | Internal Server Error | [ErrorMessage](#ErrorMessage) |

### /tasks/{id}/status

#### POST
##### Description:

update task status

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| task_id | path |  | Yes | string |
| new_status | formData |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Task successful changed status | [TaskDTO](#TaskDTO) |
| 404 | Task with provided id does not exists |  |
| 500 | Internal Server Error | [ErrorMessage](#ErrorMessage) |

### /tasks/file

#### POST
##### Description:

Create tasks from file in json or xml format

##### Responses

| Code | Description |
| ---- | ----------- |
| 201 | OK |
| 400 | Bad Request |

#### GET
##### Description:

Download dump of tasks from backend in json format

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | OK |

### Models


#### TaskDTO

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string | Unique identifier | No |
| task | string | Task string | No |
| created_at | string | UTC time in RFC3999 nano format | No |
| status | string | Status is used in sorting tasks. All new tasks have status created by default. Anyone can change any status of any task. | No |

#### UpdateTaskRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| task | string | verbose task description | No |
| status | string | Status is used in sorting tasks. All new tasks have status created by default. Anyone can change any status of any task. | No |

#### CreateTaskRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| task | string | verbose task description | Yes |
| status | string | Status is used in sorting tasks. All new tasks have status created by default. Anyone can change any status of any task. | No |

#### CreateTaskResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string | Unique identifier | No |
| task | string | Task string | No |
| created_at | string | UTC time in RFC3999 nano format | No |
| status | string | Status is used in sorting tasks. All new tasks have status created by default. Anyone can change any status of any task. | No |

#### GetTasksResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| count | integer | Total stored amount of tasks | No |
| next | string | Next page in pagination | No |
| previous | string | Previous page in pagination | No |
| result | [ [TaskDTO](#TaskDTO) ] | Resulted value of tasks | No |

#### ErrorMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| endpoint | string | full URI of requested endpoint | No |
| code | number | http code | No |
| status | string | verbose error message | No |