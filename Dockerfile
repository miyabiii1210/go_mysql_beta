FROM golang:1.19.1-bullseye

RUN apt-get update
RUN apt-get install -y vim git make ssh tree default-mysql-client

ENV PROJECT_ROOTDIR /go/app/
WORKDIR $PROJECT_ROOTDIR

COPY go/ $PROJECT_ROOTDIR/
COPY .env $PROJECT_ROOTDIR/

RUN cd $PROJECT_ROOTDIR && go mod download
RUN cd $PROJECT_ROOTDIR && go get github.com/go-sql-driver/mysql@latest github.com/joho/godotenv

EXPOSE 8080
ENV HOST 0.0.0.0