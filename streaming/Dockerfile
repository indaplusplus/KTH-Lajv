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

# Setup go auth server
RUN apt-get install golang-go git -y
RUN mkdir /stream-auth
ADD stream-auth.go /stream-auth/

ADD streaming.conf streaming.conf
WORKDIR /stream-auth
ENV GOPATH /go
RUN go get github.com/gorilla/mux
RUN go build -o auth .
WORKDIR /

CMD ["sh", "-c", "nginx -c /streaming.conf; (./stream-auth/auth &); bash"]
