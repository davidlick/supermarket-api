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

docker-build:
	docker build -t supermarket-api-image .

docker-run:
	docker run \
		--rm \
		--name supermarket-api \
		-p "3000:3000" \
		-e "ENV=dev" \
		-e "APIPORT=3000" \
		-e "LOGLEVEL=debug" \
		-e "DMLINITFILE=../../defaultproduce.json" \
		supermarket-api-image	
	
clean:
	rm ./supermarket-api ./coverage.out
