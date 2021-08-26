package server

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net/url"
	"test_Ozon_1/database"
	"test_Ozon_1/proto"
	"time"
	"unicode/utf8"
)

type GRPCServer struct{}

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	ShortUrlSize = 10
)

func isValidUrl(longURL string) (*url.URL, bool) {
	var ok bool

	u, err := url.ParseRequestURI(longURL)
	if err == nil {
		u, err = url.Parse(longURL)
		if err == nil && u.Host != "" {
			ok = true
		}
	}

	return u, ok
}

func (s *GRPCServer) Create(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	var (
		err       error
		newUrl    string
		parse_url *url.URL
		ok        bool
	)
	b := make([]byte, ShortUrlSize)
	length := utf8.RuneCountInString(letterBytes)

	if parse_url, ok = isValidUrl(req.InputUrl); ok {
		rand.Seed(time.Now().UnixNano())

		for i := range b {
			b[i] = letterBytes[rand.Intn(length)]
		}

		newUrl = parse_url.Scheme + "://" + parse_url.Host + "/" + string(b)
		err = database.InsertUrls(database.GetConn(), req.InputUrl, newUrl)

		if err != nil {
			newUrl, err = database.SearchByUrl(database.GetConn(), "long_url", req.InputUrl)
		}
	} else {
		err = errors.New("Invalid URL")
	}

	log.Printf("Create:\n\tin = %v\n\tout = %v\n", req.InputUrl, newUrl)
	return &proto.Response{TargetUrl: newUrl}, err
}

func (s *GRPCServer) Get(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	targetU, err := database.SearchByUrl(database.GetConn(), "short_url", req.GetInputUrl())

	log.Printf("Get:\n\tin = %v\n\tout = %v\n", req.InputUrl, targetU)
	return &proto.Response{TargetUrl: targetU}, err
}
