package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	tcp_ports   = []int{22, 80, 23, 445}
	timeout     = 3 * time.Second
	ip_start    = 1
	ip_end      = 254
	wg          sync.WaitGroup
	target_host string
)

func analyze(target_host string) {

	for num := 0; num < len(tcp_ports); num++ {

		address := fmt.Sprintf(target_host+":%s", strconv.Itoa(tcp_ports[num]))

		_, err := net.DialTimeout("tcp", address, timeout)
		if err == nil {
			fmt.Printf("%s\n", address)
			break
		}
	}

	wg.Done()
}

func help() {
	fmt.Println("check ports: tcp", tcp_ports)
	println("usage: ./ipscan -h localhost")
	println("usage: ./ipscan -h 192.168.11.2")
	println("usage: ./ipscan -i 192.168.11.")
}

func main() {

	if len(os.Args) < 3 {
		help()
		return
	}

	if os.Args[1] == "-h" {
		wg.Add(1)
		target_host = os.Args[2]
		analyze(target_host)
	} else if os.Args[1] == "-i" {
		for ip := ip_start; ip <= ip_end; ip++ {
			wg.Add(1)
			target_host = os.Args[2] + strconv.Itoa(ip)
			go analyze(target_host)
		}
	} else {
		help()
		return
	}
	wg.Wait()
}
