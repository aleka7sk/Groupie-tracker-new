run :
	go run cmd/main.go

docker :
	docker container rm g-t; docker image build -f dockerfile . -t imagename
	docker container run -p 9000:8000 -d --name g-t imagename
	docker ps -a
	
clean :
	docker system prune
