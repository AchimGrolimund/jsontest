# Binary name
BINARY=sos-operator

# Build values
VERSION=`git describe --tags --always`
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

tests:
	go test -v -short -race ./... # Disabled for now

build:
	@echo "Version: ${VERSION}"
	@echo "Build: ${BUILD}"
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build ${LDFLAGS} -installsuffix cgo -o ./bin/sos-operator/${BINARY}_v${VERSION} .

run:
	go run ${LDFLAGS} .

clean:
	if [ -f ./bin/sos-operator/${BINARY}_v* ] ; then rm -f ./bin/sos-operator/${BINARY}_v* ; fi

.PHONY: all