package main

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

const (
	NUMSEQ = 10
)

func main() {
	Url := "http://127.0.0.1/demo/row_locking_atomic_long_delay/"
	param := url.Values{
		"account": {"5"},
		"amount":  {"10"},
	}
	var wg sync.WaitGroup
	wg.Add(NUMSEQ)
	for i := 1; i <= NUMSEQ; i++ {
		go func(i int) {
			defer wg.Done()
			resp, err := http.PostForm(Url, param)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			fmt.Println(i, Url, resp.StatusCode)
		}(i)
	}
	wg.Wait()

}
