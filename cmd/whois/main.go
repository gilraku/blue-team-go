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
	query := flag.String("q", "", "Domain or IP to query")
	server := flag.String("s", "", "WHOIS server (auto-detected if empty)")
	flag.Parse()

	if *query == "" {
		fmt.Fprintln(os.Stderr, "Usage: whois -q <domain|ip> [-s <server>]")
		os.Exit(1)
	}

	srv := *server
	if srv == "" {
		srv = detectServer(*query)
	}

	conn, err := net.Dial("tcp", srv+":43")
	if err != nil {
		fmt.Fprintf(os.Stderr, "connection error: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s\r\n", *query)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func detectServer(q string) string {
	if net.ParseIP(q) != nil {
		return "whois.arin.net"
	}
	parts := strings.Split(q, ".")
	tld := strings.ToLower(parts[len(parts)-1])
	servers := map[string]string{
		"com": "whois.verisign-grs.com",
		"net": "whois.verisign-grs.com",
		"org": "whois.pir.org",
		"io":  "whois.nic.io",
		"id":  "whois.id",
		"uk":  "whois.nic.uk",
		"de":  "whois.denic.de",
		"ru":  "whois.tcinet.ru",
	}
	if s, ok := servers[tld]; ok {
		return s
	}
	return "whois.iana.org"
}
