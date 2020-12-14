FROM golang:1.14.3
ENV GO111MODULE "on"
ENV GOPROXY "https://goproxy.cn"
WORKDIR $GOPATH/src/todolist-backend
COPY . $GOPATH/src/todolist-backend
RUN make
EXPOSE 8080
CMD ["./main", "-c", "config/config.yaml"]
