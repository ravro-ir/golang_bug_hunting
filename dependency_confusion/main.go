package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// https://example.com/package.json <=== URL fuzzing
// supply chain attack

type PackageJson struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Main        string `json:"main"`
	Scripts     struct {
		Toc string `json:"toc"`
	} `json:"scripts"`
	Repository struct {
		Type string `json:"type"`
		Url  string `json:"url"`
	} `json:"repository"`
	Author  string `json:"author"`
	License string `json:"license"`
	Bugs    struct {
		Url string `json:"url"`
	} `json:"bugs"`
	Homepage     string `json:"homepage"`
	Dependencies struct {
		MarkdownToc string `json:"markdown-toc"`
	} `json:"dependencies"`
}

func OpenFile() string {

	jsonFile, err := os.Open("C:\\Users\\raminfp\\GolandProjects\\ravro_live_golang\\dependency_confusion\\package.json")
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
	var result PackageJson
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Name

}

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
	packageName := OpenFile()
	url := "https://www.npmjs.com/search/suggestions?q=%s"
	newUrl := fmt.Sprintf(url, packageName)
	out := HttpGet(newUrl)
	if len(out) == 2 {
		fmt.Println("[++++] Package doesn't avaliable in npm website, potential vuls deps conf")
	} else {
		fmt.Println("[----] Package does exsit")
	}
}
