FROM golang:1.17.1-alpine

LABEL maintainer="suatacikel@gmail.com"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8000

RUN go build

CMD ["./golang-rest-api"]