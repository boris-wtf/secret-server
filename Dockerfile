FROM golang:1.15.3-buster as builder

WORKDIR /go/src/github.com/KzmnbrS/boris-wtf/secret-server

RUN git clone https://github.com/boris-wtf/secret-server.git .

RUN go mod download

RUN go build -o app ./cmd/main

FROM debian:buster-20201209

WORKDIR /opt/app

COPY --from=builder /go/src/github.com/KzmnbrS/boris-wtf/secret-server/app ./app

ENTRYPOINT ./app