FROM golang:1.19.2

RUN  go mod init planScrapper
RUN go mod vendor
RUN go mod build

WORKDIR /go/src/app
COPY . .

RUN go install -v