PROJECT_ID := hew2018
SERVICE_NAME := hew2018
AWS_ID := 812656525701
AWS_REGION := us-west-2
export AWS_REGION
export AWS_DEFAULT_REGION=$(AWS_REGION)
DOCKER_IMAGE := $(AWS_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/$(PROJECT_ID):latest

deploy/build:
	env GOOS=linux env GOARCH=amd64 env CGO_ENABLED=0 go build -o ./cmd/main main.go

docker/build:
	docker build -t $(DOCKER_IMAGE) .

docker/push:
	docker push $(DOCKER_IMAGE)

aws/login:
	aws configure set region $(AWS_REGION)
	aws configure set aws_access_key_id $$AWS_ACCESS_KEY_ID
	aws configure set aws_secret_access_key $$AWS_SECRET_ACCESS_KEY

aws/deploy:
	.circleci/ecs-deploy --enable-rollback --timeout 300 --cluster $(PROJECT_ID) --service-name $(SERVICE_NAME) --image $(DOCKER_IMAGE)
