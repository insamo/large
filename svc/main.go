package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/insamo/large/proto"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
)

// TestHandler ...
type TestHandler struct {
}

// NewTestHandler ...
func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

// Update ...
func (h *TestHandler) Large(ctx context.Context, req *proto.LargeRequest, rsp *proto.LargeResponse) error {

	if b, err := ioutil.ReadFile("large.json"); err != nil {
		return err
	} else {

	rsp.Large = b
}
	return nil
}

func main() {
	svc := grpc.NewService(
		service.Name("go.micro.test"),
	)
	svc.Init()

	if err := proto.RegisterTestServiceHandler(svc.Server(), NewTestHandler()); err != nil {
		fmt.Println(err)
	}

	svc.Run()
}
