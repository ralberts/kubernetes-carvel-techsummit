ARG BASE_REGISTRY=docker.io
ARG BASE_IMAGE=golang
ARG BASE_TAG=latest

FROM golang:latest as BUILD
WORKDIR /app
COPY go.mod main.go ./
RUN \
    go install && \
    go build -o bin/simple-demo -v .

# Production
FROM ${BASE_REGISTRY}/${BASE_IMAGE}:${BASE_TAG}
WORKDIR /app
COPY --from=BUILD /app/bin/simple-demo /app/simple-demo
CMD ["./simple-demo"]