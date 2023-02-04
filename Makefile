TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=gitlab.com
NAMESPACE=kopicloud
NAME=kopicloud-ad-tf-provider
BINARY=terraform-provider-${NAME}
VERSION=$(shell git describe --tags)
BUIDL_DATE=$(shell date +%Y-%m-%d)
BUIDL_TIME=$(shell date +%T%Z)
OS=darwin
OS_ARCH=arm64
LD_FLAGAS=-ldflags '-X main.build_version=${VERSION} -X main.build_date=${BUIDL_DATE} -X main.build_time=${BUIDL_TIME}'
BUILD_CMD=go build ${LD_FLAGAS}

default: install

build:
	mkdir -p api
	go generate
	${BUILD_CMD} -o bin/${BINARY}

clean:
	rm -rf api
	find . -iname "*.gen.go" -exec rm {} \;
	rm -f ${BINARY}

release: build
	GOOS=darwin GOARCH=amd64 ${BUILD_CMD} -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 ${BUILD_CMD} -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 ${BUILD_CMD} -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm ${BUILD_CMD} -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 ${BUILD_CMD} -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 ${BUILD_CMD} -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm ${BUILD_CMD} -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 ${BUILD_CMD} -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 ${BUILD_CMD} -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 ${BUILD_CMD} -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 ${BUILD_CMD} -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 ${BUILD_CMD} -o ./bin/${BINARY}_${VERSION}_windows_amd64

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS}_${OS_ARCH}
	cp  bin/${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS}_${OS_ARCH}

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
