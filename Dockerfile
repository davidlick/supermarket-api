FROM golang:1.16-alpine

WORKDIR /go/src/github.com/davidlick/supermarket-api/
COPY . .
WORKDIR ./cmd/api/

RUN go build -o supermarket-api -v ./...

EXPOSE 3000

CMD ["./supermarket-api"]
