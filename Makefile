IMAGE_TAG ?= dev

NAME := kubectl-relay
GOBUILD = CGO_ENABLED=0 go build -trimpath

.PHONY: server-image
server-image:
	buildctl build \
		--frontend=dockerfile.v0 \
		--local context=. \
		--local dockerfile=. \
		--opt filename=manifests/Dockerfile-server \
		--output type=image,name=localhost:5001/knight42/krelay-server:$(IMAGE_TAG),push=true

.PHONY: krelay
krelay:
	$(GOBUILD) -o krelay ./cmd/client

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -race -v ./...

.PHONY: coverage
coverage:
	go test -race -v -coverprofile=cover.out ./...
	go tool cover -html cover.out
	rm cover.out

.PHONY: clean
clean:
	rm -rf krelay*
	rm -rf kubectl-relay*
