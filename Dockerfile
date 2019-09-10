FROM golang:1.13-alpine

WORKDIR /go/src/app

COPY . .

RUN go build

EXPOSE 8080
CMD ["./hello-world-go"]
