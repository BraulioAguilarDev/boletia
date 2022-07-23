FROM golang:alpine

RUN apk update && apk add git curl && \
    curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /usr/bin/air && rm -rf /var/cache/apk/*

WORKDIR /app

CMD ["air"]