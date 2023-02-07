# Go GRPC

## Generate Proto
- protoc -I proto/ proto/book.proto --go_out=plugins=grpc:v1_book_grpc/

## Before generate proto file, import
    go get -u github.com/golang/protobuf/proto
    go get -u github.com/golang/protobuf/protoc-gen-go
https://whatibroke.com/2021/05/22/protoc-gen-go-is-not-recognized-as-an-internal-or-external-command-go/
