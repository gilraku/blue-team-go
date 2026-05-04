package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	rawURL := flag.String("url", "", "URL to parse and decode")
	decode := flag.Bool("decode", false, "URL-decode the input string only")
	encode := flag.Bool("encode", false, "URL-encode the input string only")
	flag.Parse()

	if *rawURL == "" {
		fmt.Fprintln(os.Stderr, "Usage: url-parse -url <url> [-decode] [-encode]")
		os.Exit(1)
	}

	if *decode {
		out, err := url.QueryUnescape(*rawURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "decode error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(out)
		return
	}

	if *encode {
		fmt.Println(url.QueryEscape(*rawURL))
		return
	}

	u, err := url.Parse(*rawURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Scheme:    %s\n", u.Scheme)
	fmt.Printf("Host:      %s\n", u.Hostname())
	if u.Port() != "" {
		fmt.Printf("Port:      %s\n", u.Port())
	}
	fmt.Printf("Path:      %s\n", u.Path)
	if u.RawQuery != "" {
		fmt.Printf("Query:     %s\n", u.RawQuery)
		fmt.Println("\nQuery Parameters:")
		for k, v := range u.Query() {
			fmt.Printf("  %-20s = %s\n", k, strings.Join(v, ", "))
		}
	}
	if u.Fragment != "" {
		fmt.Printf("Fragment:  %s\n", u.Fragment)
	}
	if u.User != nil {
		fmt.Printf("User:      %s\n", u.User.Username())
	}

	decoded, _ := url.QueryUnescape(*rawURL)
	if decoded != *rawURL {
		fmt.Printf("\nDecoded:   %s\n", decoded)
	}
}
