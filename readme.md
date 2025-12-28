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

### Create

```
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Learn Go","status":"to-do"}'

```  

### Get

```
curl -X POST http://localhost:8080/tasks/{id} \

```

### Update

```

curl -X PUT http://localhost:8080/tasks/{id} \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go Deeply","status":"done"}'

```

### Delete

```
curl -X DELETE http://localhost:8080/tasks/{id}1

```

### List

```
curl -X GET http://localhost:8080/tasks

```

-------



