FROM golang:latest

RUN apt-get update
RUN apt-get upgrade -y

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-ozzo/ozzo-validation
RUN go get gopkg.in/mgo.v2

RUN go build -o server main/main.go
CMD ["/app/server"]

