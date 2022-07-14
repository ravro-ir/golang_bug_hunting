package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mvdan.cc/xurls/v2"
	"net/http"
	"os"
	"strings"
)

func OpenFile() string {

	jsonFile, err := os.Open("C:\\Users\\raminfp\\GolandProjects\\ravro_live_golang\\s3bucket_takeover\\README.md")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(jsonFile)
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return string(byteValue)

}

func GetXML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("data error")
	}
	return string(data), nil
}

func main() {
	const pattern = "s3.ir-thr-at1.arvanstorage.com"
	fileData := OpenFile()
	rxStrict := xurls.Strict()
	out := rxStrict.FindAllString(fileData, -1)
	for _, value := range out {
		status := strings.Contains(value, pattern)
		if status {
			xmlData, _ := GetXML(value)
			if xmlData != "" {
				checkBucket := strings.Contains(xmlData, "<Code>NoSuchBucket</Code>")
				if checkBucket {
					fmt.Println("[++++] Potential S3 Bucket Take Over")
				}

			}
		}

	}

}
