CWD=$(shell pwd)
GOPATH := $(CWD)

build:	fmt bin

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	cp -r vendor/src .

rmdeps:
	if test -d src; then rm -rf src; fi 

deps:   rmdeps
	@GOPATH=$(GOPATH) go get -u "github.com/naototty/xlsx2csv-go"
	## @GOPATH=$(GOPATH) go get -u "github.com/tealeg/xlsx"

vendor-deps: deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor/src; then rm -rf vendor/src; fi
	cp -r src vendor/src
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt main.go

bin:	self
	@GOPATH=$(GOPATH) GOOS=linux   GOARCH=amd64 go build -o xlsx2csv-go-linux64 main.go
	@GOPATH=$(GOPATH) GOOS=linux   GOARCH=amd64 go build -ldflags="-s -w" -o xlsx2csv-go-linux64-tiny main.go
	@GOPATH=$(GOPATH) GOOS=windows GOARCH=amd64 go build -o xlsx2csv-go-win64 main.go
	@GOPATH=$(GOPATH) GOOS=darwin  GOARCH=amd64 go build -o xlsx2csv-go-darwin64 main.go
