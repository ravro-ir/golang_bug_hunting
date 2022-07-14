package main

import (
	"errors"
	"fmt"
	"github.com/miekg/dns"
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
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func LookUpCNAME(domain string) ([]string, error) {
	var m dns.Msg
	var cnames []string
	m.SetQuestion(dns.Fqdn(domain), dns.TypeCNAME)
	in, err := dns.Exchange(&m, "8.8.8.8:53")
	if err != nil {
		return cnames, err
	}
	if len(in.Answer) < 1 {
		return cnames, errors.New("No Answer")
	}
	for _, answer := range in.Answer {
		if c, ok := answer.(*dns.CNAME); ok {
			cnames = append(cnames, c.Target)
		}
	}
	return cnames, nil
}

func main() {

	out, err := LookUpCNAME("testforlive.ravro.ir")
	if err != nil {
		log.Fatal(err)
	}
	for _, data := range out {
		data = strings.Replace(data, "com.", "com", -1)
		url := fmt.Sprintf("https://%s", data)
		htmlData := HttpGet(url)
		status := strings.Contains(htmlData, "There isn't a GitHub Pages site here.")
		if status {
			fmt.Println("[++++] Potential github sub domain take over")
		} else {
			fmt.Println("[----] There isn't github sub domain take over")
		}

	}

}
