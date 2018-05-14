FROM golang:1.8

ADD . /go/

WORKDIR /go/

RUN go get -d -v ./...
RUN go build -o chat .
EXPOSE 9876:9876
CMD ["/go/chat"]
