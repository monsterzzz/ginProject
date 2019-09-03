FROM golang:latest


ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

WORKDIR /Users/monster/Desktop/ginProject
COPY . /Users/monster/Desktop/ginProject
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./ginProject"]