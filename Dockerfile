FROM --platform=linux/amd64 golang:1.18 as builder

WORKDIR /src
COPY . .

RUN go build