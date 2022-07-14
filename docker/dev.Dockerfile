FROM golang:alpine

RUN apk update && \
  apk add git bash curl openssl musl-dev --no-cache

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
  && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air
