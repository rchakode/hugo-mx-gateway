PACKAGE_NAME=hugo-mx-gateway
PRODUCT_VERSION=$$(grep "ProgramVersion.=.*" main.go | cut -d"\"" -f2)-$$(git rev-parse --short HEAD)
PRODUCT_CLOUD_IMAGE_VERSION=$$(echo $(PRODUCT_VERSION) | sed 's/\.//g' -)
ARCH=$$(uname -m)
DIST_DIR=$(PACKAGE_NAME)-v$(PRODUCT_VERSION)-$(ARCH)
GOCMD=GO111MODULE=off go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -v
GOVENDOR=govendor
UPX=upx
PACKER=packer
PACKER_VERSION=1.5.1
PACKER_CONF_FILE="./deploy/packer/cloud-image.json"

all: test build

deploy:
	which gcloud
	gcloud components install app-engine-go
	gcloud app deploy --quiet

build:
	$(GOBUILD) -o $(PACKAGE_NAME) -v

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
	./$(PACKAGE_NAME)

deps: vendor

vendor:
	$(GO) mod vendor

docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v