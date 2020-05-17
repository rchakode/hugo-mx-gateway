PACKAGE_NAME=./bin/hugo-mx-gateway
DOCKER_IMAGE=rchakode/hugo-mx-gateway
ARCH=$$(uname -m)
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOVENDOR=$(GOCMD) mod vendor
GOIMAGE=golang:1.13.10
UPX=upx

all: test build

build:
	$(GOBUILD) -o $(PACKAGE_NAME) -v

build-docker:
	docker run --rm -it \
	-e GO111MODULE=on \
	-e CGO_ENABLED=0 \
	-e GOOS=linux \
	-e GOARCH=amd64 \
	-v "$(PWD)":/app \
	-w /app \
	$(GOIMAGE) \
	go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o "$(PACKAGE_NAME)" -v
	docker build -t $(DOCKER_IMAGE):$$(date +%s) .

build-compress: build
	$(UPX) $(PACKAGE_NAME)

test:
	$(GOCMD) clean -testcache
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(PACKAGE_NAME)

run:
	$(GOBUILD) -o $(PACKAGE_NAME) -v ./...
	$(PACKAGE_NAME)

deps: vendor
vendor:
	$(GOVENDOR)