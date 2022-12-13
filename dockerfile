FROM golang:1.17-alpine as base
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
WORKDIR /app

COPY . .
RUN go get github.com/codegangsta/gin
RUN go mod download
RUN go build -o spocker

EXPOSE 5005
CMD ["gin", "run", "--appPort 5001", "--port 5005", "--immediate", "--build .", "--path main.go", "--bin spocker"]