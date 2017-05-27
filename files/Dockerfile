FROM golang:1.8.3-alpine AS builder

WORKDIR /go/src/app

ADD code/* /go/src/app

RUN go-wrapper download
RUN go-wrapper install
RUN cd /go/src/app && go build -o fprocessing

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app /app/
ENTRYPOINT ./fprocessing
