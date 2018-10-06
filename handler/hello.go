package handler

import (
	"context"
	"fmt"
)

type HelloArgs struct {
	Name string
}

type HelloReply struct {
	Result string
}

type HelloSvc struct{}

func (t *HelloSvc) Hello(ctx context.Context, args *HelloArgs, reply *HelloReply) error {
	fmt.Printf("HelloSvc Hello args: %v\n", args)

	reply.Result = fmt.Sprintf("\"{\"res\":\"hello %s\"}\"", args.Name)

	fmt.Printf("HelloSvc Hello reply: %v\n", reply)

	return nil
}
