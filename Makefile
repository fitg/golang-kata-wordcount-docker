format:
	go fmt ./counter
	go fmt main.go

lint:
	golint ./counter
	go fmt main.go

vet:
	go vet ./counter
	go fmt main.go

unit-test:
	go test -v ./counter

all: format lint vet unit-test
