GOPATH:=$(shell go env GOPATH)
APP_NAME:=order-service
PROJECT_ID:=imre-experiment
IMAGE_REGISTRY:=gcr.io/imre-experiment
IMAGE_NAME:=grc.io/$(PROJECT_ID)/$(APP_NAME)
IMAGE_TAG:=latest
GIT_COMMIT:=$(shell git rev-parse --short HEAD)

.PHONY: build test docker

build: 
	CGO_ENABLED=0 go build -a -o $(APP_NAME) cmd/server/main.go

test: 
	go test -v ./... -cover -vet -all

docker: build
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

docker-build-commit: build
	docker build -t $(IMAGE_REGISTRY)/$(APP_NAME):$(GIT_COMMIT) .

docker-build: build
	docker build -t $(IMAGE_REGISTRY)/$(APP_NAME):$(IMAGE_TAG) .
	docker tag $(IMAGE_REGISTRY)/$(APP_NAME):$(IMAGE_TAG) $(IMAGE_REGISTRY)/$(APP_NAME):latest

docker-push: build
	docker push $(IMAGE_REGISTRY)/$(APP_NAME):$(IMAGE_TAG)
