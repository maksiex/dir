FROM golang:1.23.1 AS builder

WORKDIR /usr/src/backv1

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o dir_backv1 ./cmd

FROM alpine:latest

WORKDIR /usr/src/backv1

RUN apk add --no-cache ca-certificates

COPY --from=builder /usr/src/backv1/dir_backv1 .

EXPOSE 8080

COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

ENTRYPOINT ["./wait-for-it.sh", "db:5432", "./dir_backv1"]

