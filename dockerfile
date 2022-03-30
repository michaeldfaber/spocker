FROM golang:1.17-alpine as base
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
WORKDIR /app

COPY . .
RUN go mod download
# RUN go mod download github.com/gin-gonic/gin@latest
RUN go build -o spocker

# ENV GO111MODULE=off
# RUN go get -u github.com/gin-gonic/gin

EXPOSE 5005
CMD ["./gin", "--appPort 5001", "--port 5005", "--immediate", "--build .", "--path main.go", "--bin spocker"]