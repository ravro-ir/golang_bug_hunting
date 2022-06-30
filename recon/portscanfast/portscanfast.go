package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
)

func PortScan(ip string, port string, wg *sync.WaitGroup) {
	defer wg.Done()
	address := ip + ":" + port
	connection, err := net.Dial("tcp", address)
	if err == nil {
		fmt.Println("[+] Connected", connection.RemoteAddr().String())
	} else {
		fmt.Println("[-] Port " + port + " Closed")
	}
}

func main() {
	ip := flag.String("ip", "", "Please add your ip address for scanning ...")
	flag.Parse()
	flag.VisitAll(func(f *flag.Flag) {
		ipValue := f.Value.String()
		if ipValue == "" {
			log.Fatal("Error : please add ip address")
		}
	})

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
	var wg sync.WaitGroup
	for _, p := range ptrStr {
		wg.Add(1)
		go PortScan(*ip, p, &wg)
	}
	wg.Wait()
}
