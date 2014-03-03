PACKAGES := \
	arithmetic/server \
	arithmetic/client \
	arithmetic/definition
SOURCE := $(wildcard $(addsuffix /*.go, $(addprefix src/, $(PACKAGES))))
LIBRARIES := \
	github.com/entangle/goentangle
LIBRARIES_DIRS := $(addprefix src/, $(LIBRARIES))

export GOPATH=$(shell pwd)

all: arithmetic-client arithmetic-server

arithmetic-client: src/arithmetic/definition $(SOURCE) $(LIBRARIES_DIRS)
	@go build -v -o bin/artihmetic-client arithmetic/client

arithmetic-server: src/arithmetic/definition $(SOURCE) $(LIBRARIES_DIRS)
	@go build -v -o bin/artihmetic-server arithmetic/server

src/arithmetic/definition: example-definitions/arithmetic.entangle
	@entangle generate go --package definition example-definitions/arithmetic.entangle src/arithmetic/definition

$(LIBRARIES_DIRS):
	@go get $(@:src/%=%)

test: all
	@go test -v $(PACKAGES)

clean:
	@rm -rf bin pkg src/github.com src/code.google.com src/arithmetic/definition

.PHONY: test clean dist
