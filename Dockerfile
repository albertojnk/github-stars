FROM golang as builder

ENV GO111MODULE=on

WORKDIR /golang-crud-spa

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 8090

ENTRYPOINT [ "/app/golang-crud-spa" ]