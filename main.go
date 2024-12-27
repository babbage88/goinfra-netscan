package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	subnet := flag.String("subnet", "10.0.0.0/23", "subnet to be scanned")
	port := flag.Int("port", 22, "TCP port to scan for")

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", *subnet, *port), 2*time.Second)
	if err != nil {
		if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
			log.Println("Connection timed out")
		} else {
			log.Println("Connection refused", neterr)
		}
	} else {
		log.Println("Connection successful to", conn.RemoteAddr().String())
		conn.Close()
	}
}
