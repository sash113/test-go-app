FROM golang:1.9-alpine

RUN apk add --no-cache ca-certificates \
        dpkg \
        gcc \
        git \
        musl-dev \
        bash

ADD .docker/go/entrypoint.sh /entrypoint.sh

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" && chmod +x /entrypoint.sh
RUN go get github.com/tockins/realize

WORKDIR $GOPATH/src

ENTRYPOINT ["/entrypoint.sh"]
CMD ["realize", "start"]