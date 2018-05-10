FROM golang:1.8

ADD database.go /go/database.go
WORKDIR /go/

RUN go get -d -v ./...
RUN go build -o db .
RUN touch kthlive.db
EXPOSE 55994:55994
CMD ["/go/db"]
