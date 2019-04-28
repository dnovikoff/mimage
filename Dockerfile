FROM golang:latest as builder

ADD . /go/src/github.com/dnovikoff/mimage
WORKDIR /go/src/github.com/dnovikoff/mimage
RUN VERBOSE=1 CGO_ENABLED=0 make all

FROM scratch
COPY --from=builder /go/src/github.com/dnovikoff/mimage/gobin/mimage /usr/bin/mimage
EXPOSE 8080
COPY pkg/image/test_data/sprite.png /etc/mimage/sprite.png

ENTRYPOINT ["/usr/bin/mimage", "--sprite", "/etc/mimage/sprite.png"]