FROM ubuntu:16.04

MAINTAINER Viktor Edström

# Compiling nginx + the rtmp module takes ages
# Probably better of just downloading the dep from somewhere.
ADD debs.tar.gz ./

RUN apt-get update -y
RUN dpkg -i --force-all ./debs/*.deb; exit 0
RUN apt-get install -f -y
# Expose RTMP
EXPOSE 1935:1935
# Auth works at 1337

# Web
EXPOSE 1339:1339
# HLS
#EXPOSE 8080:8080



# Setup go auth server
RUN apt-get install golang-go git -y
RUN mkdir /stream-auth
ADD stream-auth.go /stream-auth/
ADD stream.go /stream-auth/
ADD DatabaseConnector.go /stream-auth/

ADD streaming.conf streaming.conf
WORKDIR /stream-auth
ENV GOPATH /go
RUN go get github.com/gorilla/mux
RUN go build -o auth stream-auth.go DatabaseConnector.go
RUN go build -o web stream.go DatabaseConnector.go

WORKDIR /
RUN mkdir /tmp/hls
RUN chmod 755 /tmp/hls
CMD ["sh", "-c", "nginx -c /streaming.conf; (./stream-auth/auth &); ./stream-auth/web"]
