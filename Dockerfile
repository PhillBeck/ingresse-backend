FROM golang:1.10-alpine

WORKDIR /go/src/github.com/PhillBeck/ingresse-backend

RUN apk update && apk add git && apk add tzdata \
    && cp -r -f /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime

COPY . .

RUN go get -v github.com/golang/mock/gomock/...

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 5000

CMD ["ingresse-backend"]