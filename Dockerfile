# Docker multi stage build
## Builds executable
FROM golang:1.12-stretch as builder

RUN mkdir -p $GOPATH/src/github.com/r0nni3
WORKDIR $GOPATH/src/github.com/r0nni3
ENV GO111MODULE=on
ADD . ./backend-test
RUN go get -v ...
RUN cd backend-test && go test -v ./... && go build -o /bin/app


# Final image where app runs
FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

COPY --from=builder /bin/app ./import

CMD exec /bin/import
