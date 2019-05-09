#!/bin/bash

cd /go/src/${GO_PROJECT_DIR}

go get -d ./...

exec "$@"
