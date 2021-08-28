package main

import (
	grpc "google.golang.org/grpc"
	"log"
	"net"
	"os"
	"shortener/conf"
	"shortener/proto"
	"shortener/server"
)

func main() {
    log.Println("------------------- Server start -------------------")
    
	database_ip, ok := os.LookupEnv("DB_HOST")
	if ok {
		conf.DB_conf.Host = database_ip
		log.Printf("New database host has been set. New host = %s", database_ip)
	} else {
		log.Printf("Database host = %s", conf.DB_conf.Host)
	}

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
    
    log.Println("------------------- Server stop -------------------")
}
