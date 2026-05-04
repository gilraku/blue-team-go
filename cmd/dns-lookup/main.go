package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	host := flag.String("host", "", "Hostname or IP to query")
	rtype := flag.String("type", "A", "Record type: A, AAAA, MX, NS, TXT, CNAME, PTR")
	flag.Parse()

	if *host == "" {
		fmt.Fprintln(os.Stderr, "Usage: dns-lookup -host <hostname> [-type <record>]")
		os.Exit(1)
	}

	switch strings.ToUpper(*rtype) {
	case "A":
		addrs, err := net.LookupHost(*host)
		if err != nil {
			fatal(err)
		}
		for _, a := range addrs {
			if !strings.Contains(a, ":") {
				fmt.Println(a)
			}
		}
	case "AAAA":
		addrs, err := net.LookupHost(*host)
		if err != nil {
			fatal(err)
		}
		for _, a := range addrs {
			if strings.Contains(a, ":") {
				fmt.Println(a)
			}
		}
	case "MX":
		records, err := net.LookupMX(*host)
		if err != nil {
			fatal(err)
		}
		for _, r := range records {
			fmt.Printf("%d\t%s\n", r.Pref, r.Host)
		}
	case "NS":
		records, err := net.LookupNS(*host)
		if err != nil {
			fatal(err)
		}
		for _, r := range records {
			fmt.Println(r.Host)
		}
	case "TXT":
		records, err := net.LookupTXT(*host)
		if err != nil {
			fatal(err)
		}
		for _, r := range records {
			fmt.Println(r)
		}
	case "CNAME":
		cname, err := net.LookupCNAME(*host)
		if err != nil {
			fatal(err)
		}
		fmt.Println(cname)
	case "PTR":
		names, err := net.LookupAddr(*host)
		if err != nil {
			fatal(err)
		}
		for _, n := range names {
			fmt.Println(n)
		}
	default:
		fmt.Fprintf(os.Stderr, "unsupported record type: %s\n", *rtype)
		os.Exit(1)
	}
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "lookup error: %v\n", err)
	os.Exit(1)
}
