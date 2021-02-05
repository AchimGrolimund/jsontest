# Binary name
BINARY=sos-operator

# Build values
VERSION=`git describe --tags --always`
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

tests: buildHtml_min
	go test -v -short -race ./... # Disabled for now

build: buildHtml_min
	@echo "Version: ${VERSION}"
	@echo "Build: ${BUILD}"
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build ${LDFLAGS} -installsuffix cgo -o ./bin/sos-operator/${BINARY}_v${VERSION} .

run: buildHtml_min
	go run ${LDFLAGS} .

clean:
	if [ -f ./bin/sos-operator/${BINARY}_v* ] ; then rm -f ./bin/sos-operator/${BINARY}_v* ; fi

buildHtml_min:
	tr -d '\r\n\t' < email.html > email.min.html.tmp && tr -s " " < email.min.html.tmp > email.min.html && rm -r email.min.html.tmp

.PHONY: all