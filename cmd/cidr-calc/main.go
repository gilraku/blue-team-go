package main

import (
	"flag"
	"fmt"
	"math/big"
	"net"
	"os"
)

func main() {
	cidr := flag.String("cidr", "", "CIDR notation, e.g. 192.168.1.0/24")
	flag.Parse()

	if *cidr == "" {
		fmt.Fprintln(os.Stderr, "Usage: cidr-calc -cidr <network>")
		os.Exit(1)
	}

	ip, network, err := net.ParseCIDR(*cidr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid CIDR: %v\n", err)
		os.Exit(1)
	}

	ones, bits := network.Mask.Size()
	hostBits := bits - ones
	total := new(big.Int).Lsh(big.NewInt(1), uint(hostBits))

	firstIP := cloneIP(network.IP)
	lastIP := cloneIP(network.IP)
	for i := range lastIP {
		lastIP[i] |= ^network.Mask[i]
	}

	fmt.Printf("Input IP:       %s\n", ip)
	fmt.Printf("Network:        %s\n", network.IP)
	fmt.Printf("Broadcast:      %s\n", lastIP)
	fmt.Printf("Subnet Mask:    %s\n", net.IP(network.Mask))
	fmt.Printf("Prefix Length:  /%d\n", ones)
	fmt.Printf("Total Hosts:    %s\n", total)
	fmt.Printf("Usable Hosts:   %s\n", usable(total))
	fmt.Printf("First Usable:   %s\n", nextIP(firstIP))
	fmt.Printf("Last Usable:    %s\n", prevIP(lastIP))
}

func cloneIP(ip net.IP) net.IP {
	clone := make(net.IP, len(ip))
	copy(clone, ip)
	return clone
}

func nextIP(ip net.IP) net.IP {
	next := cloneIP(ip)
	for i := len(next) - 1; i >= 0; i-- {
		next[i]++
		if next[i] != 0 {
			break
		}
	}
	return next
}

func prevIP(ip net.IP) net.IP {
	prev := cloneIP(ip)
	for i := len(prev) - 1; i >= 0; i-- {
		prev[i]--
		if prev[i] != 255 {
			break
		}
	}
	return prev
}

func usable(total *big.Int) *big.Int {
	if total.Cmp(big.NewInt(2)) <= 0 {
		return big.NewInt(0)
	}
	return new(big.Int).Sub(total, big.NewInt(2))
}
