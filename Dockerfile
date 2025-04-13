# Build Stage
FROM golang:1.20 AS build

WORKDIR /go/src/wizapp
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/src/wizapp/wizapp

# Release Stage
FROM alpine:3.17

# Required for Go binary execution
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy built binary
COPY --from=build /go/src/wizapp/wizapp .

# Copy any required assets or files
COPY wizexercise.txt /app/wizexercise.txt

EXPOSE 8080

# Execute the binary directly
ENTRYPOINT ["/app/wizapp"]



