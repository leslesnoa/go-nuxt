FROM golang:1.17.2-alpine as builder

WORKDIR /go/api

COPY . ./

ENV GO111MODULE=on

RUN CGO_ENABLE=0 GOOS=linux go build -a installsuffix cgo .

FROM alpine:latest as prod

EXPOSE 8080

WORKDIR /api

COPY --from=builder /go/api/ .

RUN pwd

CMD ["./webapp-api"]

FROM golang:1.17.2-alpine as dev

EXPOSE 8080

WORKDIR /go/src/whoisapp

COPY . ./

ENV GO111MODULE=on

RUN go get github.com/pilu/fresh

CMD ["fresh"]