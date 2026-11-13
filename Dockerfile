# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Cache module downloads
COPY go.mod go.sum ./
RUN go mod download

# Build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/efilm-api .

# Runtime stage
FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /out/efilm-api /app/efilm-api

EXPOSE 8080
ENV APP_PORT=8080

ENTRYPOINT ["/app/efilm-api"]
