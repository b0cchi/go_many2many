FROM golang:1.20.0

RUN apt update && apt install && \
    go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/spf13/cobra-cli@latest && \
    go install golang.org/x/tools/cmd/godoc@latest

WORKDIR /go/src