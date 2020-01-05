

ifdef CIRCLE_SHA1
TAG ?= $(CIRCLE_SHA1)
else
TAG ?= $(shell git rev-parse HEAD)
endif

PROJECT_NAME ?= tara
PROJECT_DOCKER_HOST ?= zengzhiyuan



run:
	go run main.go

# docker

build:
	docker build -t $(PROJECT_NAME):$(TAG) .
