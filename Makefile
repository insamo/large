
GOPATH:=$(shell go env GOPATH)

INCLUDES := -I=.
 
GOGOOUT:="\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:."


.PHONY: proto
proto:
	protoc ${INCLUDES} --proto_path=${GOPATH}/src:. --micro_out=. --go_out=plugins=micro:. proto/*.proto

.PHONY: service
service:
	go build -o service svc/main.go

.PHONY: client
client:
	go build -o client cli/main.go

.PHONY: all
all: proto service client

.DEFAULT_GOAL := all