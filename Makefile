format:
	go fmt ./...
	cd counter
	go fmt

lint:
	golint ./...
	golint ./counter

vet:
	go vet ./...
	cd counter
	go vet

build:
	docker build . -t golang-kata-wordcount-docker

unit-test:
	docker run --name tester golang-kata-wordcount-docker sh -c "cd ./counter && go test -v ./..."
	docker rm tester

run:
	docker run --name runner golang-kata-wordcount-docker sh -c "./main"
	docker rm runner

all: format lint vet build unit-test run
