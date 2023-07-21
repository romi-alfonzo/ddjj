FROM golang:1.20.6-alpine

WORKDIR /usr/src/app

RUN apk update \
    && apk add poppler-dev poppler poppler-utils \
    && pdftotext -v \
    && rm -rf /var/cache/apk/*

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY parser/go.mod parser/go.sum ./
RUN go mod download && go mod verify

COPY parser .
RUN go build -v -o /usr/local/bin/ ./...

CMD ["parser", "serve"]
