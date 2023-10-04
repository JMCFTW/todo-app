build: Dockerfile
	docker build . -t todo-app
run:
	docker run -d --rm -p 8080:8080 -it todo-app
stop:
	docker ps -a | grep todo-app | awk '{print $$1}' | xargs docker stop