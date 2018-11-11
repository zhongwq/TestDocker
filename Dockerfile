FROM golang:1.8
MAINTAINER WilsonZhong "1316628630@qq.com"
WORKDIR $GOPATH/src/github.com/zhongwq/TestDocker
ADD . $GOPATH/src/github.com/zhongwq/TestDocker
RUN go get github.com/codegangsta/negroni
RUN go get github.com/gorilla/mux
RUN go get github.com/unrolled/render
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/spf13/pflag
RUN go build .
EXPOSE 9999
ENTRYPOINT ["./TestDocker"]
