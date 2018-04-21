package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{
		Dial: dial,
	},
}

func get(target, keyword string, p *plugin) {
	target = strings.Replace(target, "{}", keyword, 1)
	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:56.0) Gecko/20100101 Firefox/56.0")
	req.Header.Set("Referer", p.Information.WebSite)
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	check(resp, p)
}

func post(target, keyword string, p *plugin) {
	u := url.Values{}

	// 添加 json 里配置的请求
	for k, v := range p.Request.PostFields {
		if v == "" {
			u.Set(k, keyword)
		}
	}

	body := ioutil.NopCloser(strings.NewReader(u.Encode()))

	req, err := http.NewRequest(http.MethodPost, target, body)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:56.0) Gecko/20100101 Firefox/56.0")
	req.Header.Set("Referer", p.Information.WebSite)
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	check(resp, p)
}

// dial 设置超时
func dial(netw, addr string) (net.Conn, error) {
	deadline := time.Now().Add(3 * time.Second)
	c, err := net.DialTimeout(netw, addr, time.Second*10)
	if err != nil {
		return nil, err
	}

	c.SetDeadline(deadline)
	return c, nil
}
