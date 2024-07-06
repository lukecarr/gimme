FROM gitpod/workspace-go:latest

ARG GO_VERSION="1.22.5"

# install latest go
RUN go install golang.org/dl/go$GO_VERSION@latest && \
    go$GO_VERSION download && \
    echo "alias go=go$GO_VERSION" > ~/.bashrc \
