FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY *.go ./
COPY .env ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o lema .

FROM scratch

WORKDIR /root/

COPY --from=builder /app/lema .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./lema", "server"]