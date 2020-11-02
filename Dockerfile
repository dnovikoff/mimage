FROM golang:1.15 as builder

ADD . /build
WORKDIR /build
RUN GOPROXY=off VERBOSE=1 CGO_ENABLED=0 make all

FROM scratch
COPY --from=builder /build/gobin/mimage /usr/bin/mimage
EXPOSE 8080
COPY ./pkg/image/test_data/sprite.png /etc/mimage/sprite.png

ENTRYPOINT ["/usr/bin/mimage", "--sprite", "/etc/mimage/sprite.png"]