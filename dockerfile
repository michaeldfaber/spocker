FROM golang:1.17-alpine as base
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
WORKDIR /app

COPY . .
RUN go get github.com/codegangsta/gin
RUN go mod download
RUN go build -o spocker
RUN chmod 777 scripts/run.sh

EXPOSE 5005
ENV HOST=0.0.0.0

ENTRYPOINT ["scripts/run.sh"]