FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY ./cmd/paneful_scraper/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o paneful_scraper
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/paneful_scraper .
CMD ["./paneful_scraper"]