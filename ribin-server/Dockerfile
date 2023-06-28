FROM alpine:3.14

RUN mkdir /app
WORKDIR /app

COPY ./game-server ./
COPY ./conf.yaml ./

CMD ["./game-server"]
