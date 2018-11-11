FROM golang:1.8
COPY . "$GOPATH/src/github.com/zhongwq/"
RUN cd "$GOPATH/src/github.com/zhongwq/TestDocker" && go get -v && go install -v
WORKDIR /
EXPOSE 9999
VOLUME ["/data"]