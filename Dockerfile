FROM golang:1.15.3-alpine3.12 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main .

CMD  ["/app/main", "serve"]