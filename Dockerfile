FROM golang:1.11

WORKDIR /go/src/github.com/picfight/pfcd
COPY . .

RUN env GO111MODULE=on go install . ./cmd/...

EXPOSE 9108

CMD [ "pfcd" ]
