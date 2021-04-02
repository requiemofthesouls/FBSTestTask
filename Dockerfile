FROM golang:1.16.2 AS builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/app/main.go

FROM ubuntu:latest
WORKDIR /root/
COPY --from=builder /src .

EXPOSE 8080
EXPOSE 5300
ENTRYPOINT ["./app"]
