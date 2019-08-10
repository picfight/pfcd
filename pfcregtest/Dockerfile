FROM golang:1.11

WORKDIR /go/src/github.com/picfight/pfcregtest
COPY . .

RUN apt-get update && apt-get upgrade -y && apt-get install -y rsync

RUN git clone https://github.com/picfight/pfcd /go/src/github.com/picfight/pfcd

RUN cd /go/src/github.com/picfight/pfcd && env GO111MODULE=on go install . .\cmd\...

RUN pfcd --version
