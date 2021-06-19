dev:
	docker-compose up

unit-test:
	go test -short ./...

integration-test:
	go test -integration ./...

all-test:
	go test ./...

test-coverage:
	go test -cover ./...

cover-report:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

build:
	go build -o supermarket-api cmd/api/*.go

run:
	go run cmd/api/*.go

clean:
	rm ./supermarket-api ./coverage.out
