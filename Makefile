VERSION:=$(shell cat VERSION)
REGISTRY_DOMAIN=ghcr.io
REGISTRY_NAME=ghcr.io/LEMARBE/Fortune

test:
	go test -timeout=5s -count=1 -p 4 -race -covermode=atomic -coverprofile=coverage.out ./...

linux:
	env GOOS=linux GOARCH=amd64 \
	go build \
	-buildvcs=false \
	-ldflags "-s -X 'main.Version=$(VERSION)'" \
	-o build/linux

mac:
	env GOOS=darwin GOARCH=amd64 \
	go build \
	-buildvcs=false \
	-ldflags "-s -X 'main.Version=$(VERSION)'" \
	-o build/mac

.PHONY: build
build: export TAG=$(VERSION)
build:
	docker build --no-cache -f ./Dockerfile -t ${REGISTRY_NAME}:${TAG} .
	docker tag ${REGISTRY_NAME}:${TAG} ${REGISTRY_NAME}:latest

push: export TAG=$(VERSION)
push:
	docker push ${REGISTRY_NAME}:${TAG}
	docker push ${REGISTRY_NAME}:latest

generate:
	go generate ./...
