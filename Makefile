BINARY_NAME=server

swag-v1: ### swag init
	./swag init -g internal/controller/http/v1/router.go
.PHONY: swag-v1

run: swag-v1 ### swag run
	go mod tidy && go mod download && \
	GIN_MODE=debug CGO_ENABLED=0 go run -tags migrate ./cmd/app --c="config/config.yml"
.PHONY: run

build: swag-v1 
	go mod tidy && go mod download && \
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux ./cmd/app/main.go
