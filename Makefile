GOPATH := $(shell pwd)
.PHONY: clean test

all:
	@GOPATH=$(GOPATH) go install qtunnel

clean:
	@rm -fr bin pkg

test:
	@GOPATH=$(GOPATH) go test tunnel
