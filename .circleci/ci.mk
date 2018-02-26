PROJECT-ID := deploy-test-191702
REGION := us-central1-a
APP := main
DOCKER_IMAGE := gcr.io/$(PROJECT-ID)/$(APP)

login:
	gcloud config set project $(PROJECT-ID)
	gcloud config set compute/zone $(REGION)

go/build:
	make create
	env GOOS=linux env GOARCH=amd64 env CGO_ENABLED=0 go build -o ./cmd/main main.go

docker/build:
	docker build -t $(DOCKER_IMAGE) .

docker/deploy:
	docker tag $(DOCKER_IMAGE) $(DOCKER_IMAGE):$(TIMESTAMP)
	gcloud docker -- push $(DOCKER_IMAGE):$(TIMESTAMP)
	kubectl set image deployment/$(APP) $(APP)=$(DOCKER_IMAGE):$(TIMESTAMP)

init:
	#gcloud container clusters create $(APP)-cluster --machine-type f1-micro --disk-size=30 --num-nodes=3
	gcloud container clusters resize $(APP)-cluster --size=1 -q

first:
	kubectl run $(APP) --image=$(DOCKER_IMAGE) --replicas=2 --port=8000
	kubectl expose deployment $(APP) --type=LoadBalancer --port 80 --target-port 8000