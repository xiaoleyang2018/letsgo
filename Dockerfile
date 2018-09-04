FROM golang AS app
RUN go version

COPY . /go/src/github.com/qinhao/letsgo

WORKDIR /go/src/github.com/qinhao/letsgo

RUN go build -a -o app .
# RUN /bin/sh

ENTRYPOINT ["./app"]