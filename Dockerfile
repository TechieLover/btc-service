# Take base image as golang
FROM golang:1.10

# Copy Source code to work space

ADD . /go/src/vdart/btc-service

RUN go install vdart/btc-service

ENTRYPOINT /go/bin/btc-service

#Service listens on port 3000.
EXPOSE 3000
