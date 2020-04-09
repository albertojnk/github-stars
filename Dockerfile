FROM golang:latest

ENV GO111MODULE=on MONGO_URI="mongodb://mongodb:27018" ES_URL="http://elasticsearch:9200"

WORKDIR /app

COPY ./backend ./backend
COPY ./main.go .

RUN go mod init github-stars
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

COPY ./frontend ./frontend

RUN apt-get install curl && curl -sL https://deb.nodesource.com/setup_10.x | bash - && apt-get install nodejs

RUN cd frontend/app && rm -rf node_modules && rm -rf dist

RUN cd frontend/app && npm install && npm run build

ENTRYPOINT [ "/app/github-stars" ]