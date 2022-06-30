package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("data error")
	}
	return data, nil
}

func main() {
	const url = "https://s3.ir-thr-at1.arvanstorage.com/"

	var wordlist = []string{
		"google",
		"qwe",
		"ada",
	}
	for _, word := range wordlist {
		newUrl := url + word
		fmt.Println("[+++] URL is : ", newUrl)
		if xmlByte, err := GetXML(newUrl); err != nil {
			log.Printf("Failed to get xml %v", err)
		} else {
			fmt.Println(string(xmlByte))
		}
	}
}
