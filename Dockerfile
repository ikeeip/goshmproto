FROM golang:1.9.2-alpine3.6 as builder

RUN apk --no-cache add --update \
    make \
    g++ \
    linux-headers \
    git \
    glide

ENV APP_PATH=/go/src/github.com/ikeeip/goshmproto
ADD . $APP_PATH
WORKDIR $APP_PATH
RUN glide install \
 && GOOS=linux GOARCH=amd64 go build -v -x -ldflags '-extldflags "-static"' -o bin/probe .

FROM alpine:3.6

COPY --from=builder /go/src/github.com/ikeeip/goshmproto/bin/probe /probe
EXPOSE 8889
CMD ["/probe"]
