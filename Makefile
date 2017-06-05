# build everything

.DEFAULT_GOAL := all
BINDIR := bin

build_go:
	glide install
	mkdir -p $(BINDIR)
	go build -o bin/server cmd/server/main.go
	go build -o bin/importer cmd/importer/main.go

build_client:
	cd client && npm install && npm run build

all: build_go build_client

test:
	go test $$(glide novendor)

clean:
	rm -rf bin
	rm -rf vendor
	rm -rf client/node_modules
	rm -rf client/build
	glide cc
