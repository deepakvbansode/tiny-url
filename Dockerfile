FROM golang:1.15.2-alpine AS builder
RUN apk update && apk add  build-base autoconf automake libtool g++ make

COPY . /app

ENV GOARCH=amd64
ENV GOOS=linux
ENV CGO_ENABLED=0


WORKDIR /app

RUN go mod download
RUN make build

FROM alpine

# Below 2 lines should be removed if passed from env variables
RUN mkdir -p /etc/config
COPY /etc/config /etc/config

COPY --from=builder /app/bin/application /bin/application

ENTRYPOINT ["/bin/application"]
