package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"os"
	"shortener/proto"
)

func main() {
	server_ip, _ := os.LookupEnv("SERVER_HOST")
	server_ip += ":8080"
	log.Printf("Server ip = %s", server_ip)
	log.Println("Client start")
	conn, err := grpc.Dial(server_ip, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewShortenerClient(conn)

	longU := getCmdParam()
	res, err := client.Create(context.Background(), &proto.Request{InputUrl: longU})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("After Create\nPassed in 'Create': %v\nResult 'Create': %v\n", longU, res.TargetUrl)
	shortU := res.GetTargetUrl()

	res, err = client.Get(context.Background(), &proto.Request{InputUrl: shortU})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("After Get:\nPassed in 'Get': %v\nResult 'Get': %v\n", shortU, res.TargetUrl)
}

func getCmdParam() string {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("not enough args")
	}

	return flag.Arg(0)
}
