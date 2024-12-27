package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	prettyPrinter "github.com/babbage88/goinfra-netscan/internal/pretty"
)

func main() {
	pretty := prettyPrinter.NewPrettyPrinter()
	subnet := flag.String("subnet", "10.0.0.0/23", "subnet to be scanned")
	port := flag.Int("port", 22, "TCP port to scan for")
	timeout := flag.Int("timeout", 5, "Number of Second for timeout")
	flag.Parse()

	ips, err := parseCIDRstr(*&subnet)
	if err != nil {
		pretty.PrintErrorf("Error parsing provided subnet: %s", *subnet)
		pretty.PrintErrorf("Error: ", err)
	}

	for _, ip := range ips {

		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, *port), time.Duration(*timeout)*time.Second)
		if err != nil {
			if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				pretty.PrintWarning("Connection timed out")
			} else {
				pretty.PrintErrorf("Connection refused", neterr)
			}
		} else {
			pretty.Printf("Connection successful to", conn.RemoteAddr().String())
			conn.Close()
		}
	}
}
