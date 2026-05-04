package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

var securityHeaders = []string{
	"Strict-Transport-Security",
	"Content-Security-Policy",
	"X-Frame-Options",
	"X-Content-Type-Options",
	"Referrer-Policy",
	"Permissions-Policy",
	"X-XSS-Protection",
}

func main() {
	url := flag.String("url", "", "URL to inspect")
	insecure := flag.Bool("k", false, "Skip TLS verification")
	flag.Parse()

	if *url == "" {
		fmt.Fprintln(os.Stderr, "Usage: http-inspect -url <url> [-k]")
		os.Exit(1)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: *insecure},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(*url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Printf("URL:     %s\n", *url)
	fmt.Printf("Status:  %s\n", resp.Status)
	fmt.Printf("Proto:   %s\n\n", resp.Proto)

	fmt.Println("=== ALL HEADERS ===")
	var keys []string
	for k := range resp.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("  %-40s %s\n", k+":", strings.Join(resp.Header[k], "; "))
	}

	fmt.Println("\n=== SECURITY HEADERS ===")
	for _, h := range securityHeaders {
		v := resp.Header.Get(h)
		if v != "" {
			fmt.Printf("  [✓] %-35s %s\n", h, v)
		} else {
			fmt.Printf("  [✗] %s (missing)\n", h)
		}
	}
}
