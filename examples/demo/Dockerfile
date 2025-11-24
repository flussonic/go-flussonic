FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY go.mod ./
COPY go.sum* ./
RUN go mod download

# Copy source code
COPY . .

# Build the demo application
WORKDIR /build/examples/demo
RUN go build -o demo main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /build/examples/demo/demo /app/demo
COPY --from=builder /build/examples/demo/streams.csv /app/streams.csv

CMD ["/app/demo"]

