package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	ip := flag.String("ip", "", "IP address to look up ASN for")
	flag.Parse()

	if *ip == "" {
		fmt.Fprintln(os.Stderr, "Usage: asn-lookup -ip <address>")
		os.Exit(1)
	}

	if net.ParseIP(*ip) == nil {
		fmt.Fprintln(os.Stderr, "invalid IP address")
		os.Exit(1)
	}

	// Team Cymru IP-to-ASN whois service
	conn, err := net.Dial("tcp", "whois.cymru.com:43")
	if err != nil {
		fmt.Fprintf(os.Stderr, "connection error: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "begin\nverbose\n%s\nend\n", *ip)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "AS") || strings.HasPrefix(line, "Bulk") {
			continue
		}
		if line != "" {
			fields := strings.SplitN(line, "|", 5)
			if len(fields) == 5 {
				fmt.Printf("ASN:     %s\n", strings.TrimSpace(fields[0]))
				fmt.Printf("Prefix:  %s\n", strings.TrimSpace(fields[1]))
				fmt.Printf("Country: %s\n", strings.TrimSpace(fields[2]))
				fmt.Printf("RIR:     %s\n", strings.TrimSpace(fields[3]))
				fmt.Printf("ISP:     %s\n", strings.TrimSpace(fields[4]))
			}
		}
	}
}
