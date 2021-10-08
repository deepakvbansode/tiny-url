FROM golang:alpine AS builder
COPY . /app

ENV GOARCH=amd64
ENV GOOS=linux
ENV CGO_ENABLED=0


ARG GITLAB_ID
ARG GITLAB_TOKEN

WORKDIR /app

RUN go mod download
RUN make build

FROM alpine

COPY --from=builder /app/bin/application /bin/application

ENTRYPOINT ["/bin/application"]
