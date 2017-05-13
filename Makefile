# build everything

.DEFAULT_GOAL := all

build_go:
	glide install
	go build -o server cmd/server/main.go
	go build -o importer cmd/importer/main.go

build_ui:
	cd ui && npm install && npm run build

all: build_go build_ui

test:
	go test $$(glide novendor)

clean:
	$(RM) server importer
	rm -rf vendor
	rm -rf ui/node_modules
	rm -rf ui/build
	glide cc
