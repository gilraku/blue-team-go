package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := flag.String("r", "", "IP range: CIDR (10.0.0.0/24), range (10.0.0.1-10), or list (10.0.0.1,10.0.0.2)")
	count := flag.Bool("count", false, "Only print count, not individual IPs")
	flag.Parse()

	if *input == "" {
		fmt.Fprintln(os.Stderr, "Usage: ip-range -r <cidr|range|list> [-count]")
		os.Exit(1)
	}

	var ips []string

	switch {
	case strings.Contains(*input, "/"):
		ip, network, err := net.ParseCIDR(*input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid CIDR: %v\n", err)
			os.Exit(1)
		}
		_ = ip
		for ip2 := cloneIP(network.IP); network.Contains(ip2); inc(ip2) {
			ips = append(ips, ip2.String())
		}
	case strings.Contains(*input, "-"):
		parts := strings.SplitN(*input, "-", 2)
		baseIP := net.ParseIP(parts[0])
		if baseIP == nil {
			fmt.Fprintf(os.Stderr, "invalid IP: %s\n", parts[0])
			os.Exit(1)
		}
		baseIP = baseIP.To4()
		endOctet, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid range end: %s\n", parts[1])
			os.Exit(1)
		}
		startOctet := int(baseIP[3])
		for i := startOctet; i <= endOctet; i++ {
			ip2 := cloneIP(baseIP)
			ip2[3] = byte(i)
			ips = append(ips, ip2.String())
		}
	default:
		ips = strings.Split(*input, ",")
		for i, ip := range ips {
			ips[i] = strings.TrimSpace(ip)
		}
	}

	if *count {
		fmt.Printf("%d\n", len(ips))
		return
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func cloneIP(ip net.IP) net.IP {
	c := make(net.IP, len(ip))
	copy(c, ip)
	return c
}

func inc(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}
