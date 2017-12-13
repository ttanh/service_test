FROM golang:1.9

RUN apt-get update

RUN go get -v github.com/jinzhu/gorm && \
    go get -v github.com/go-sql-driver/mysql && \
    go get -v github.com/sirupsen/logrus && \
    go get -v github.com/gin-gonic/gin && \
    go get -v gopkg.in/asaskevich/govalidator.v4

RUN mkdir /go/src/service_test

EXPOSE 9000