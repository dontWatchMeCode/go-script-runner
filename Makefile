dev:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go

build:
	go build -o ./main main.go

docker:
	docker run \
		-w /app  \
		-v $(shell pwd):/app  \
		golang:1.20.4-alpine3.17 \
		sh -c "go build -o main main.go && chown $(shell id -u):$(shell id -g) main"
