FROM golang:1.20.5-alpine3.18 as builder

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY ./jiro-stamp-rally/go.mod ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]