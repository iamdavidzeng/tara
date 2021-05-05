.PHONY: test build
ifdef CIRCLE_SHA1
TAG ?= $(CIRCLE_SHA1)
else
TAG ?= $(shell git rev-parse HEAD)
endif

PROJECT_NAME ?= tara
REMOTE_DOCKER_HOST ?= zengzhiyuan


run:
	go run cmd/tara/main.go

test:
	go test -v -race ./...

# docker

build:
	docker build -f build/Dockerfile -t $(PROJECT_NAME):$(TAG) .

docker-login:
	echo $$DOCKER_PASSWORD | docker login --username=$(DOCKER_USERNAME) --password-stdin

docker-save:
	mkdir -p docker-images
	docker save -o docker-images/docker_images.tar $(PROJECT_NAME):$(TAG)

docker-load:
	docker load -i docker-images/docker_images.tar

docker-tag:
	docker tag $(PROJECT_NAME):$(TAG) $(REMOTE_DOCKER_HOST)/$(PROJECT_NAME):$(TAG)

push-images:
	docker push $(REMOTE_DOCKER_HOST)/$(PROJECT_NAME):$(TAG) 

pull-images:
	docker pull $(REMOTE_DOCKER_HOST)/$(PROJECT_NAME):$(TAG)
