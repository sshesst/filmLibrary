FROM golang:latest AS builder

WORKDIR /filmLibrary

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o filmLibrary cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /filmLibrary/filmLibrary .

COPY config/server.toml .

CMD ["./filmLibrary"]
