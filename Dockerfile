# syntax=docker/dockerfile:1
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG TARGET=web
ARG PROJECT_NAME=chi-go-template
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/${PROJECT_NAME} ./cmd/${TARGET}
RUN mkdir -p /app/storage/logs

FROM alpine:3.21
WORKDIR /app
ARG PROJECT_NAME=chi-go-template
ARG TARGET=web
COPY --from=builder /app/storage /app/storage
COPY --from=builder /out/${PROJECT_NAME} /app/${PROJECT_NAME}
ENV BINARY=${PROJECT_NAME}
EXPOSE 8080
ENTRYPOINT ["/bin/sh", "-c", "exec /app/${BINARY}"]
