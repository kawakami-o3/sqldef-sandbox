FROM golang:1.16.3-alpine3.12
RUN apk add --no-cache build-base make gcc git

COPY . /sqldef
WORKDIR /sqldef
RUN export GOPATH=/go/; \
  export GOBIN=$HOME/bin; \
  make all && sh -ec "mv build/*/mysqldef /usr/bin/ && mv build/*/psqldef /usr/bin/ && mv build/*/sqlite3def /usr/bin/"
