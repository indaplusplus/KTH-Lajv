FROM golang:1.8

ADD . /go/

WORKDIR /go/

RUN go get -d -v ./...
RUN go build -o comments .
EXPOSE 8080:8080
CMD ["/go/comments"]
