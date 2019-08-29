FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
COPY ./backend ./backend
COPY ./main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 8090

ENTRYPOINT [ "/app/golang-crud-spa" ]