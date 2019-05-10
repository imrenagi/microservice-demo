GOPATH:=$(shell go env GOPATH)
APP_NAME:=order-service
PROJECT_ID:=imre-experiment
IMAGE_REGISTRY:=gcr.io/imre-experiment
IMAGE_NAME:=grc.io/$(PROJECT_ID)/$(APP_NAME)
IMAGE_TAG:=latest
GIT_COMMIT:=$(shell git rev-parse --short HEAD)

.PHONY: build test docker

build-order: 
	CGO_ENABLED=0 go build -a -o order order-service/cmd/server/main.go

build-payment: 
	CGO_ENABLED=0 go build -a -o payment payment-service/cmd/server/main.go

test: 
	go test -v ./... -cover -vet -all

docker: 
	docker build -t $(IMAGE_REGISTRY)/order-service:v1.0.3 -f Dockerfile-Order .
	docker build -t $(IMAGE_REGISTRY)/payment-service:v1.0.4 -f Dockerfile-Payment .

docker-push:
	docker push $(IMAGE_REGISTRY)/order-service:v1.0.3
	docker push $(IMAGE_REGISTRY)/payment-service:v1.0.4
