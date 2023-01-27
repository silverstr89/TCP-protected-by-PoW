# base go image
FROM golang:1.18-alpine as  client

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o client .

RUN chmod +x /app/server

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=client /app/client /app

EXPOSE 8080/tcp

CMD [ "/app/client" ]