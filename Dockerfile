FROM golang:1.11

WORKDIR /go/src/github.com/picfight/pfcd
COPY . .

RUN apt-get update && apt-get upgrade -y && apt-get install -y rsync
RUN env GO111MODULE=on go install . ./cmd/...

EXPOSE 9108
EXPOSE 19108

CMD [ "pfcd" ]
