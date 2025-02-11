package main

import (
	"net"

	"github.com/babbage88/goinfra-netscan/internal/pretty"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func parseCIDRstr(subnet string) ([]string, error) {
	var ips []string
	ip, ipnet, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, err
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		pretty.Printf("%s", ip.String())
		ips = append(ips, ip.String())
	}
	numIps := len(ips)
	// Checking if a /32 was was parsed, in which case dont try removing broadcast netId
	switch numIps {
	case 1:
		return ips, nil
	default:
		ips = ips[1 : numIps-1]
	}
	// remove network address and broadcast address
	return ips, nil
}
