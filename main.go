package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/babbage88/goinfra-netscan/internal/pretty"
)

func scanFunc(iport string, timeout time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp", iport, timeout)
	if err != nil {
		if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
			pretty.PrintWarning("Connection timed out")
		} else {
			pretty.PrintError("Connection refused", neterr.Error())
		}
	} else {
		pretty.Print("Connection successful to", conn.RemoteAddr().String())
		conn.Close()
	}
}

func main() {
	subnet := flag.String("subnet", "10.0.0.0/23", "subnet to be scanned")
	port := flag.Int("port", 22, "TCP port to scan for")
	timeout := flag.Int("timeout", 5, "Number of Second for timeout")
	flag.Parse()

	ips, err := parseCIDRstr(*subnet)
	if err != nil {
		pretty.PrintErrorf("Error parsing provided subnet: %s", *subnet)
		pretty.PrintError("Error: ", err)
	}
	timeoutSec := time.Duration(*timeout) * time.Second
	var wg sync.WaitGroup

	for _, ip := range ips {
		iport := fmt.Sprintf("%s:%d", ip, *port)
		wg.Add(1) // Increment the counter before starting a goroutine
		go scanFunc(iport, timeoutSec, &wg)
	}

	wg.Wait()
}
