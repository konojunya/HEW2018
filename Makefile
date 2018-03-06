create:
	cp config.yml.template config.yml
	
migrate:
	go run migration/main.go

seed: migrate
	go run seed/main.go

go/build:
	make -f .circleci/ci.mk go/build

docker/build: go/build
	make -f .circleci/ci.mk docker/build

docker/push: go/build
	env NODE_ENV=production npm run build
	make -f .circleci/ci.mk docker/build
	make -f .circleci/ci.mk login
	make -f .circleci/ci.mk docker/push

docker/seed/build:
	make -f .circleci/ci.mk docker/seed/build

docker/seed/push:
	make -f .circleci/ci.mk docker/seed/push

docker/seed/deploy:
	make -f .circleci/ci.mk docker/seed/deploy