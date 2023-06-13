FROM golang:1.19

ENV GO111MODULE=on

WORKDIR /app

RUN mkdir /var/log/entry/
RUN go mod init newfeed_app
RUN go mod download
COPY ../../ .

# RUN go build cmd/newsfeed_svc/main.go

# CMD ["/app/main -config config/files/test.yml"]
CMD ["go run main.go"]