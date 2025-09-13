FROM public.ecr.aws/docker/library/golang:1.25.0-alpine AS builder

WORKDIR /go/src/github.com/yuita-yoshihiko/go-sample-api
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git alpine-sdk
RUN set -x \
    && go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /server ./cmd/main.go
