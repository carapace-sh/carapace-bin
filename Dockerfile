FROM ubuntu

RUN apt-get update \
 && apt-get install -y gccgo-go

RUN apt-get install -y git upx

ENV GOPATH /go
