package main

import (
	"database/sql"
	grpc "google.golang.org/grpc"
	"log"
	"net"
	"test_Ozon_1/proto"
	"test_Ozon_1/server"
)

var (
	db *sql.DB
)

func main() {
	s := grpc.NewServer()
	srv := &server.GRPCServer{}
	proto.RegisterShortenerServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
