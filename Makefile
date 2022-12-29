TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=gitlab.com
NAMESPACE=kopicloud
NAME=kopicloud-ad-tf-provider
BINARY=${NAME}
VERSION=0.0.1
OS=darwin
OS_ARCH=arm64

default: install

build:
	export PATH=$GOPATH/bin:$PATH
	$(MAKE) -C generator build
	$(MAKE) -C terraform build

clean:
	export PATH=$GOPATH/bin:$PATH
	$(MAKE) -C generator clean
	$(MAKE) -C terraform clean

install: build
	export PATH=$GOPATH/bin:$PATH
	$(MAKE) -C generator install
	$(MAKE) -C terraform install

