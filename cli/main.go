package main

import (
	"context"
	"fmt"
	"time"

	"github.com/insamo/large/proto"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {

	c := grpc.NewClient(
		grpc.MaxSendMsgSize(1024*1024*8),
	)

	s := proto.NewTestService("go.micro.test", c)
	callOpts := []client.CallOption{
		client.WithRequestTimeout(1 * time.Minute),
		client.WithDialTimeout(5 * time.Second),
		client.WithRetries(1),
	}

	if resp, err := s.Large(context.TODO(), &proto.LargeRequest{}, callOpts...); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(resp)
	}
}
