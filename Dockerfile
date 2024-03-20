FROM golang:latest

RUN go version

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o film-library ./cmd/main.go

CMD ["./film-library"]