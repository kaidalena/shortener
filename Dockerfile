FROM golang:1.17

RUN mkdir -p $GOPATH/src/github.com/shortener
ADD . $GOPATH/src/github.com/shortener

WORKDIR $GOPATH/src/github.com/shortener

RUN go get -u github.com/golang/protobuf@v1.5.0
RUN go get -u google.golang.org/grpc@v1.40.0
RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go get -u github.com/lib/pq@v1.10.2
RUN go mod tidy
