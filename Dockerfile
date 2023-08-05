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

RUN go build $DIR/cmd/api/main.go


##    S E R V E
## ********************************************
FROM alpine

ENV GIN_MODE=release
ENV SERVICE_NAME=id_manager
ENV PORT=80

COPY --from=builder /go/src/app/main /bin/id-manager

COPY docker/docker-entrypoint.sh /usr/local/bin/
ENTRYPOINT ["docker-entrypoint.sh"]

CMD ["/bin/id-manager"]
EXPOSE 80
