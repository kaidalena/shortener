package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"test_Ozon_1/proto"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewShortenerClient(conn)

	longU := getCmdParam()
	res, err := client.Create(context.Background(), &proto.Request{InputUrl: longU})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nPassed in 'Create': %v\nResult 'Create': %v\n", longU, res.TargetUrl)
	shortU := res.GetTargetUrl()

	res, err = client.Get(context.Background(), &proto.Request{InputUrl: shortU})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nPassed in 'Get': %v\nResult 'Get': %v\n", shortU, res.TargetUrl)
}

func getCmdParam() string {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("not enough args")
	}

	return flag.Arg(0)
}
