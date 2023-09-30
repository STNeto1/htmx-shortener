FROM golang:1.21.1-alpine3.18 as builder

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /app

ARG PORT
ARG BASE_URL
ARG DATABASE_URL="file:./db.sqlite3"

COPY go.mod go.sum ./
RUN go mod tidy 

COPY . .

RUN go build -o app .

FROM alpine:3.18.3 as production


COPY --from=builder app .

EXPOSE 3000 

CMD ./app
