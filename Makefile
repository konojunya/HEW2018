HOST := registry.heroku.com
APP_NAME := hew2018
PLATFORM := web
DOCKER_IMAGE := $(HOST)/$(APP_NAME)/$(PLATFORM)

go/build:
	env GOOS=linux env GOARCH=amd64 env CGO_ENABLED=0 go build -o ./cmd/main main.go

docker/login:
	docker login -u $(DOCKER_USER) -p $(DOCKER_PASS) $(HOST)

docker/build:
	docker build -t $(DOCKER_IMAGE) .

docker/push:
	docker push $(DOCKER_IMAGE)
