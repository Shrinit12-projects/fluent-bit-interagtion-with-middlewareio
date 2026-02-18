FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -o app

RUN mkdir -p /var/log

CMD ["./app"]
