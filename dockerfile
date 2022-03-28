FROM golang:1.17-alpine as base
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o spocker
EXPOSE 5001
CMD ["go", "run", "."]