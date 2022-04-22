# Step 1, build the binary
FROM golang:1.18.1-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 go build -o /webserver

# Step 2, copy the binary into a small image
FROM scratch

COPY --from=builder /webserver /webserver

EXPOSE 8080

ENTRYPOINT ["/webserver"]
