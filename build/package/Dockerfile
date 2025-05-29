FROM golang:1.24-alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o paneful_scraper
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/paneful_scraper .
CMD ["./paneful_scraper"]