package main

import (
	"net"
	"sync"
	"time"

	"github.com/babbage88/goinfra-netscan/internal/pretty"
)

// Function to scan for a given IPAddress:PortNumber
func PortScan(iport string, timeout time.Duration, wg *sync.WaitGroup) (bool, error) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp", iport, timeout)
	if err != nil {
		if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
			pretty.PrintWarning("Connection timed out")
			conn.Close()
			return false, neterr
		} else {
			pretty.PrintError("Connection refused", neterr.Error())
			conn.Close()
			return false, neterr
		}
	} else {
		pretty.Print("Connection successful to", conn.RemoteAddr().String())
		conn.Close()
		return true, nil
	}
}

// Function to scan for a given IPAddress:PortNumber
func PortScanfunc(iport string, timeout time.Duration, wg *sync.WaitGroup, onlyActive bool) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp", iport, timeout)
	if err != nil {
		if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
			switch onlyActive {
			case false:
				pretty.PrintWarning("Connection timed out")
			}
		} else {
			switch onlyActive {
			case false:
				pretty.PrintError("Connection refused", neterr.Error())
			}
		}
		return // Ensure we exit the function early if there's an error
	}
	defer conn.Close() // Close the connection only if it was successfully established

	pretty.Print("Connection successful to", conn.RemoteAddr().String())
}
