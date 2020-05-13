FROM golang

RUN mkdir -p /go/src/github.com/egawata/gohello
COPY . /go/src/github.com/egawata/gohello/
RUN go install github.com/egawata/gohello

EXPOSE 8088
ENTRYPOINT ["/go/bin/gohello"]
