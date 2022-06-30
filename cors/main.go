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
	if response.Header.Get("access-control-allow-origin") == "*" {
		fmt.Println("[+++] This site potential CORS vuls.")
	} else {
		fmt.Println("[+++] This site is not CORS vuls.")
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
