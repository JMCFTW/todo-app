Build and run it on localhost:8080 with Docker.
```shell
make build && make run
```

Stop run it on localhost:8080
```shell
make stop
```

Get todos:
```shell
curl http://localhost:8080/api/todos
```
Add todo:
```shell
curl -X POST http://localhost:8080/api/todos -H "Content-Type: application/json" -d '{"title": "Jimmy", "done": false, "body": "do domething"}'
```

Patch todo:
```shell
curl -X PATCH http://localhost:8080/api/todos/1/done
```