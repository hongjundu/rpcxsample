package handler

import (
	"context"
	"fmt"
)

type AddArgs struct {
	Left  int
	Right int
}

type AddReply struct {
	Result int
}

type AddSvc struct{}

func (t *AddSvc) Add(ctx context.Context, args *AddArgs, reply *AddReply) error {
	fmt.Printf("AddSvc Add args: %v\n", args)

	reply.Result = args.Left + args.Right

	fmt.Printf("AddSvc Add reply: %v\n", reply)

	return nil
}
