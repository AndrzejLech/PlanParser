FROM golang:1.19.2

RUN go get github.com/gin-gonic/gin

WORKDIR /go/src/app
COPY . .

RUN go install -v