FROM golang:1.22
WORKDIR /go/src/github.com/lukecarr/gimme/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM gcr.io/distroless/static
COPY --from=0 /go/src/github.com/lukecarr/gimme/gimme .
ENTRYPOINT ["/gimme"]