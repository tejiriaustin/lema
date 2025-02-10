FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev sqlite-dev

COPY go.mod .
COPY *.go ./
COPY .env ./

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o lema .

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache sqlite

COPY --from=builder /app/lema .
COPY --from=builder /app/.env .

VOLUME ["/app"]

EXPOSE 8080

CMD ["./lema", "api"]
