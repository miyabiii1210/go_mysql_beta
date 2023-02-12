FROM golang:1.19.1-bullseye

ENV PROJECT_ROOTDIR ./app/
WORKDIR $PROJECT_ROOTDIR
RUN mkdir ./go
COPY go/ ./go
RUN cd ./go && go mod download
RUN cd ./go; go get github.com/go-sql-driver/mysql@latest github.com/joho/godotenv

EXPOSE 8080
ENV HOST 0.0.0.0