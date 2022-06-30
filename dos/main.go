package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

const BASE_URL = ""
const NUM_PARALLEL = 1000

type result struct {
	bodyStr string
	err     error
}

func HttpGet() (string, error) {

	req, err := http.NewRequest("GET", BASE_URL, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	bodyStr := string(body)
	return bodyStr, nil

}

func AsyncHttp(data []string) ([]string, error) {
	var wg sync.WaitGroup
	wg.Add(NUM_PARALLEL)
	resultCh := make(chan result)
	for i := 0; i < NUM_PARALLEL; i++ {
		go func() {
			for input := range data {
				fmt.Println("Counter request is : ", input)
				bodyStr, err := HttpGet()
				resultCh <- result{bodyStr, err}

			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()
	results := []string{}
	for reslt := range resultCh {
		if reslt.err != nil {
			return nil, reslt.err
		}
		results = append(results, reslt.bodyStr)
	}

	return results, nil
}

func main() {

	data := []string{}
	for i := 1; i <= 100000; i++ {
		data = append(data, strconv.Itoa(i))
	}
	myresult, err := AsyncHttp(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, res := range myresult {
		fmt.Println(res)
	}
}
