package server

import (
	"context"
	"net/url"
	"regexp"
	"shortener/proto"
	"testing"
)

var (
	shortUrls []string
)

func TestCreate(t *testing.T) {
	testTable := []struct {
		url      string
		expected bool
	}{
		{
			url:      "http://test.org/one",
			expected: true,
		},
		{
			url:      "http://test.org/one",
			expected: true,
		},
		{
			url:      "https://red.ru/two/three",
			expected: true,
		},
		{
			url:      "green.ru/fufuf",
			expected: false,
		},
		{
			url:      "http://blue.com",
			expected: true,
		},
		{
			url:      "",
			expected: false,
		},
	}

	server := GRPCServer{}
	for _, testCase := range testTable {
		result, err := server.Create(context.Background(), &proto.Request{InputUrl: testCase.url})
		if testCase.expected {
			if err != nil || result.TargetUrl == "" {
				t.Errorf("Wrong answer: in = (%s), out = (%s, %v)", testCase.url, result.TargetUrl, err)
			} else {
				if !checkHostUrls(testCase.url, result.TargetUrl) || !checkShortUrl(result.TargetUrl) {
					t.Errorf("+ Wrong answer: in = (%s), out = (%s, %v).", testCase.url, result.TargetUrl, err)
				}
				shortUrls = append(shortUrls, result.TargetUrl)
			}
		} else {
			if err == nil {
				t.Errorf("Wrong answer: in = (%s), out = (%s, %v)", testCase.url, result.TargetUrl, err)
			}
		}
	}
}

func TestGet(t *testing.T) {
	testTable := []struct {
		url      string
		expected bool
	}{
		{
			url:      "1test short url",
			expected: false,
		},
		{
			url:      "2test short url",
			expected: false,
		},
	}

	for _, nextUrl := range shortUrls {
		testTable = append(testTable, struct {
			url      string
			expected bool
		}{url: nextUrl, expected: true})
	}

	server := GRPCServer{}
	for _, testCase := range testTable {
		result, err := server.Get(context.Background(), &proto.Request{InputUrl: testCase.url})
		if testCase.expected {
			if err != nil || result.TargetUrl == "" {
				t.Errorf("Wrong answer: in = (%s), out = (%s, %v)", testCase.url, result.TargetUrl, err)
			} else {
				if !checkHostUrls(testCase.url, result.TargetUrl) {
					t.Errorf("Wrong answer: in = (%s), out = (%s, %v). Hosts are not equal.", testCase.url, result.TargetUrl, err)
				}
			}
		} else {
			if err == nil {
				t.Errorf("Wrong answer: in = (%s), out = (%s, %v)", testCase.url, result.TargetUrl, err)
			}
		}
	}
}

func checkHostUrls(inUrl, outUrl string) bool {
	inUParse, _ := url.ParseRequestURI(inUrl)
	outUParsed, err2 := url.ParseRequestURI(outUrl)
	return err2 == nil && inUParse.Host == outUParsed.Host
}

func checkShortUrl(shortUrl string) bool {
	parseU, err := url.ParseRequestURI(shortUrl)
	if err != nil {
		return false
	}

	m, err := regexp.MatchString("^/[a-zA-Z0-9_]{10}$", parseU.Path)
	return m && err == nil
}
