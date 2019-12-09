FROM golang:1.13-alpine3.10 as builder

RUN apk --no-cache --no-progress add make git

WORKDIR /go/src/github.com/AubreyHewes/ledgo
COPY . .
RUN make build

FROM alpine:3.10
RUN apk update \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates

COPY --from=builder /go/src/github.com/AubreyHewes/ledgo/dist/ledgo /usr/bin/ledgo
ENTRYPOINT [ "/usr/bin/ledgo" ]
