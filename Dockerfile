FROM golang:1.15.3

ADD . /go/src/app
WORKDIR /go/src/app

ENV PORT 8080
EXPOSE 8080

RUN apk add --no-cache git

RUN go get

RUN apk del git

RUN go build -o main .

CMD ["/go/src/app/main"]