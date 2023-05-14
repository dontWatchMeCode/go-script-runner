build:
	go build -o ./main main.go

build-docker:
	docker run \
		-w /app  \
		-v $(shell pwd):/app  \
		-v $(shell pwd)/.cache/gocache:/go/cache \
		-v $(shell pwd)/.cache/gopath:/go/pkg/mod \
		golang:1.20.4-alpine3.17 \
		sh -c "go build -o main main.go && \
			chown $(shell id -u):$(shell id -g) main && \
			chown -R $(shell id -u):$(shell id -g) .cache"

dev:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go

dev-run:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go -run
