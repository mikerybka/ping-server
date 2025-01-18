FROM alpine:latest

RUN apk update
RUN apk add go

RUN go install github.com/mikerybka/ping-server@latest

ENTRYPOINT ["/root/go/bin/ping-server"]
