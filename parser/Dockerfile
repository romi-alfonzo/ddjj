FROM golang:1.14-stretch

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN apt-get update && apt-get install poppler-utils -y

RUN go build -mod=readonly -o server

CMD [ "./server" ]
