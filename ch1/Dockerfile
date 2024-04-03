# syntax=docker/dockerfile:1
FROM golang:1.21

WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN go build -o /hello-world

CMD ["/hello-world"]

