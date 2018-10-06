package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"log"
	"mysamples/rpcxsample/handler"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	consulAddr = flag.String("consulAddr", "localhost:8500", "server address")
	basePath   = flag.String("base", "/rpcx", "prefix path")
)

func main() {
	flag.Parse()

	helloD := client.NewConsulDiscovery(*basePath, "HelloSvc", []string{*consulAddr}, nil)
	helloClient := client.NewXClient("HelloSvc", client.Failtry, client.RandomSelect, helloD, client.DefaultOption)
	defer helloClient.Close()

	addD := client.NewConsulDiscovery(*basePath, "AddSvc", []string{*consulAddr}, nil)
	addClient := client.NewXClient("AddSvc", client.Failtry, client.RandomSelect, addD, client.DefaultOption)
	defer addClient.Close()

	go func() {

		for {
			args := &handler.HelloArgs{
				Name: "Du Hongjun",
			}

			reply := &handler.HelloReply{}
			err := helloClient.Call(context.Background(), "Hello", args, reply)
			if err != nil {
				log.Printf("[ERROR] failed to call Hello: %v", err)
				continue
			}

			log.Printf("Hello Reply: %v", reply.Result)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			args := &handler.AddArgs{
				Left:  1,
				Right: 2,
			}

			reply := &handler.AddReply{}
			err := addClient.Call(context.Background(), "Add", args, reply)
			if err != nil {
				log.Printf("[ERROR] failed to call Add: %v", err)
				continue
			}

			log.Printf("Add Reply: %v", reply.Result)
			time.Sleep(time.Second)
		}
	}()

	exitchan := make(chan os.Signal, 1)
	signal.Notify(exitchan, os.Interrupt)
	signal.Notify(exitchan, syscall.SIGTERM)
	signal.Notify(exitchan, syscall.SIGKILL)
	<-exitchan

}
