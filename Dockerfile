##    P A C K E R
## ********************************************
FROM golang:1.20-alpine AS packer

ENV DIR=/go/src/app
WORKDIR $DIR

RUN apk add curl

COPY app/ $DIR/
COPY docker/*-entrypoint.sh /usr/local/bin/

RUN go mod download

ENTRYPOINT ["docker-entrypoint.sh"]


##    B U I L D E R
## ********************************************
FROM packer AS builder

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN go build


##    S E R V E
## ********************************************
FROM alpine

ENV GRPC_PORT=9090

COPY --from=builder /go/src/app/id-manager /bin/id-manager

COPY docker/docker-entrypoint.sh /usr/local/bin/
ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["/bin/id-manager"]
EXPOSE 9090
