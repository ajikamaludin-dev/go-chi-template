# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG TARGET=web
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/app ./cmd/${TARGET}

FROM gcr.io/distroless/static-debian12
COPY --from=builder /out/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
