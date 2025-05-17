.DEFAULT_GOAL := run

CURRENT_DIR := $(shell pwd)
APP_CMD_DIR := ${CURRENT_DIR}/cmd
APP := $(shell basename ${CURRENT_DIR})

TAG := latest
ENV_TAG := latest
PROJECT_NAME := ${PROJECT_NAME}

build:
	go mod download && GOOS=linux go build -tags musl -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/${APP}/main.go

clear:
	rm -rf ${CURRENT_DIR}/bin/*

network:
	docker network create --driver=bridge ${NETWORK_NAME}

build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

run:
	go run ${APP_CMD_DIR}/${APP}/main.go
