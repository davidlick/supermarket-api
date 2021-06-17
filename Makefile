dev:
	docker-compose up

unit-test:
	go test -short ./...

integration-test:
	go test -integration ./...

all-test:
	go test ./...

build:
	go build -o supermarket-api cmd/api/*.go

run:
	go run cmd/api/*.go

clean:
	rm ./supermarket-api
