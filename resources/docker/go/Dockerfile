FROM golang:1.7.5

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

# this will ideally be built by the ONBUILD below ;)
CMD ["go-wrapper", "run"]

ADD . /go/src/app
ONBUILD RUN go-wrapper download
ONBUILD RUN go-wrapper install
ONBUILD RUN buffalo build
