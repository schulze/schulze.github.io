package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

func main() {
	proxyURL, err := url.Parse("socks5://127.0.0.1:9050")
	dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{Dial: dialer.Dial}
	client := http.Client{Transport: transport}

	resp, err := client.Get("https://check.torproject.org")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
