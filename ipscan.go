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
	tcp_ports       = []int{22, 80, 23}
	timeout         = 3 * time.Second
	ip_start        = 1
	ip_end          = 254
	wg              sync.WaitGroup
	target_host     string
	supplement_port int
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
	println("usage: ipscan [-o|--one] [hostname|ip] [-p|--port port_number]")
	println("       ipscan [-m|--multi] 192.168.11. [-p|--port port_number]")
	println("       ipscan -h")
}

func main() {

	if len(os.Args) < 3 || len(os.Args) == 4 {
		help()
		return
	} else if len(os.Args) >= 5 && (os.Args[3] == "-p" || os.Args[3] == "--port") {
		for p := 4; p < len(os.Args); p++ {
			supplement_port, _ = strconv.Atoi(os.Args[p])
			if supplement_port < 1 || 65534 < supplement_port {
				println("[quit] port_number between 1 and 65534. error:", os.Args[p])
				return
			}
			tcp_ports = append(tcp_ports, supplement_port)
		}
	}

	if os.Args[1] == "-h" {
		help()
		return
	} else if os.Args[1] == "-o" || os.Args[1] == "--one" {
		wg.Add(1)
		target_host = os.Args[2]
		analyze(target_host)
	} else if os.Args[1] == "-m" || os.Args[1] == "--multi" {
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
