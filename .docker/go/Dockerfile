FROM golang:1.9-alpine

RUN apk add --no-cache ca-certificates \
        dpkg \
        gcc \
        git \
        musl-dev \
        bash

ADD src/ /go/src/${GO_PROJECT_DIR}
ADD .docker/go/entrypoint.sh /entrypoint.sh

WORKDIR /go/src/${GO_PROJECT_DIR}

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" && chmod +x /entrypoint.sh
RUN go get github.com/pilu/fresh
RUN go get -d
# RUN go build

ENTRYPOINT ["/entrypoint.sh"]

CMD /go/bin/fresh
