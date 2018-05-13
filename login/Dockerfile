FROM golang:1.8

ADD Database.go /go/Database.go
ADD Login.go /go/Login.go
ADD Self.go /go/Self.go

WORKDIR /go/

RUN go get -d -v ./...
RUN go build -o login .
EXPOSE 8021:8021
ENV LOGIN_API_KEY CHANGEME
CMD ["/go/login"]
