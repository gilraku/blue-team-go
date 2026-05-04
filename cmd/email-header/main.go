package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
)

var reReceived = regexp.MustCompile(`from\s+(\S+)\s+\(.*?\[(\d+\.\d+\.\d+\.\d+)\]`)
var reSpamScore = regexp.MustCompile(`(?i)X-Spam-Score:\s*([0-9.\-]+)`)

func main() {
	file := flag.String("file", "", "Raw email file to parse (reads stdin if empty)")
	flag.Parse()

	var scanner *bufio.Scanner
	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		scanner = bufio.NewScanner(f)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	headers := map[string][]string{}
	var order []string
	var current string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
			if current != "" {
				headers[current][len(headers[current])-1] += " " + strings.TrimSpace(line)
			}
		} else {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				k := strings.TrimSpace(parts[0])
				v := strings.TrimSpace(parts[1])
				if _, exists := headers[k]; !exists {
					order = append(order, k)
				}
				headers[k] = append(headers[k], v)
				current = k
			}
		}
	}

	printField := func(label, key string) {
		if v, ok := headers[key]; ok {
			fmt.Printf("%-18s %s\n", label+":", v[0])
		}
	}

	fmt.Println("=== KEY HEADERS ===")
	printField("From", "From")
	printField("To", "To")
	printField("Subject", "Subject")
	printField("Date", "Date")
	printField("Message-ID", "Message-ID")
	printField("Reply-To", "Reply-To")
	printField("Return-Path", "Return-Path")

	fmt.Println("\n=== AUTHENTICATION ===")
	for _, k := range []string{"Authentication-Results", "DKIM-Signature", "ARC-Authentication-Results"} {
		if v, ok := headers[k]; ok {
			fmt.Printf("%-18s %s\n", k+":", v[0])
		}
	}
	printField("SPF", "Received-SPF")

	if v, ok := headers["X-Spam-Score"]; ok {
		fmt.Printf("\n%-18s %s\n", "Spam Score:", v[0])
	}

	fmt.Println("\n=== HOP ANALYSIS ===")
	if hops, ok := headers["Received"]; ok {
		for i, hop := range hops {
			fmt.Printf("Hop %d: %s\n", i+1, hop[:min2(len(hop), 100)])
			if m := reReceived.FindStringSubmatch(hop); m != nil {
				rdns, _ := net.LookupAddr(m[2])
				rstr := "no rDNS"
				if len(rdns) > 0 {
					rstr = rdns[0]
				}
				fmt.Printf("       IP: %s  rDNS: %s\n", m[2], rstr)
			}
		}
	}
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
