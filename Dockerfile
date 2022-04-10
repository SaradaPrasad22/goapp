# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /goapp

EXPOSE 8080
RUN chmod -R 777  /goapp
CMD [ "/goapp" ]
