FROM golang:1.9-alpine

WORKDIR /go/src/github.com/PhillBeck/ingresse-backend

RUN apk update && apk add git && apk add tzdata \
    && cp -r -f /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime

COPY . .
RUN go-wrapper download
RUN go-wrapper install

EXPOSE 5000

CMD ["go-wrapper", "run"]