package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

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
		"21", "22", "25", "80", "2100", "8080", "8090", "8990",
	}
	for _, p := range ports {
		address := *ip + ":" + p
		connection, err := net.Dial("tcp", address)
		if err == nil {
			fmt.Println("[+] Connected", connection.RemoteAddr().String())
		} else {
			fmt.Println("[-] Port " + p + " Closed")
		}
	}
}
