FROM golang:1.18.4-alpine3.16

WORKDIR /usr/src/app

COPY .env.example ./.env
COPY config.yml.example ./config.yml

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/server *.go

EXPOSE 8088

CMD ["server"]