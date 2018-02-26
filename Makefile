HOST := registry.heroku.com
APP_NAME := hew2018
PLATFORM := web
DOCKER_IMAGE := $(HOST)/$(APP_NAME)/$(PLATFORM)

local/start:
	make migrate
	make go/build
	npm install
	npm run build
	docker build -t hew2018 -f dev.dockerfile .
	docker run -p 8000:8000 hew2018

migrate:
	go run cmd/migration.go

go/build:
	env GOOS=linux env GOARCH=amd64 env CGO_ENABLED=0 go build -o ./cmd/main main.go

docker/login:
	docker login -u $(DOCKER_USER) -p $(DOCKER_PASS) $(HOST)

docker/build:
	docker build -t $(DOCKER_IMAGE) .

docker/push:
	docker push $(DOCKER_IMAGE)
