# 📝 Task App – Go REST API

A clean, layered **Task Management REST API** written in **Go**, following **industry best practices** such as Repository–Service–Handler architecture, dependency injection, and concurrency-safe in-memory storage.

## 📌 Features

- CRUD operations for Tasks
- Clean separation of concerns
- Thread-safe in-memory repository
- RESTful HTTP API
- Easy to extend to DB-backed storage
- Idiomatic Go code

---

## APIs

### Create Task

```
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Learn Go","status":"to-do"}'

```  

### Get Task

```
curl -X POST http://localhost:8080/tasks/{id} \

```

### Update Task

```

curl -X PUT http://localhost:8080/tasks/{id} \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go Deeply","status":"done"}'

```

### Delete Task

```
curl -X DELETE http://localhost:8080/tasks/{id}1

```

### List Tasks

```
curl -X GET http://localhost:8080/tasks

```

-------


## Docker Build 

   <pre>docker build -t  task_app . </pre>

## Docker Run

   <pre>docker run -p 8080:8080 task_app:latest  </pre> 



