package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func HttpGet(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := new(http.Client)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if response.Header.Get("x-frame-options") == "" {
		fmt.Println("[+++] This site potential clickjacking vuls.")
	} else {
		fmt.Println("[+++] This site is not clickjacking vuls.")
	}
}

func main() {
	domain := flag.String("domain", "", "Please add your domain address for scanning ...")
	flag.Parse()
	flag.VisitAll(func(f *flag.Flag) {
		ipValue := f.Value.String()
		if ipValue == "" {
			log.Fatal("Error : please add domain address")
		}
	})
	HttpGet(*domain)

}
