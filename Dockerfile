FROM golang:latest as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/github.com/ak1ra24/slack-grafana-image-renderer-picker
COPY . .
RUN go build  ./cmd/gfslack/gfslack.go

FROM frolvlad/alpine-glibc

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/ak1ra24/slack-grafana-image-renderer-picker/gfslack /usr/local/bin/gfslack

ENTRYPOINT ["/usr/local/bin/gfslack"]
