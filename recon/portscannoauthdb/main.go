package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
)

// MongoDb
// Redis
// Memcatch
// jenkins
// hood Hbase
// Elasticsearch
// Casandra db

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
	ports := []string{
		"27017", "6379", "11211", "5984", "8080", "80", "2003", "50470", "8020", "9000", "9200",
	}

	var wg sync.WaitGroup
	for _, port := range ports {
		wg.Add(1)
		go PortScan(*ip, port, &wg)
	}
	wg.Wait()
}
