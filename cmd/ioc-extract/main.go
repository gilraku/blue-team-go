package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
)

var (
	reIPv4   = regexp.MustCompile(`\b(?:25[0-5]|2[0-4]\d|[01]?\d\d?)(?:\.(?:25[0-5]|2[0-4]\d|[01]?\d\d?)){3}\b`)
	reDomain = regexp.MustCompile(`\b(?:[a-zA-Z0-9](?:[a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}\b`)
	reMD5    = regexp.MustCompile(`\b[0-9a-fA-F]{32}\b`)
	reSHA1   = regexp.MustCompile(`\b[0-9a-fA-F]{40}\b`)
	reSHA256 = regexp.MustCompile(`\b[0-9a-fA-F]{64}\b`)
	reURL    = regexp.MustCompile(`https?://[^\s"'<>]+`)
	reEmail  = regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
)

func main() {
	file := flag.String("file", "", "File to scan (reads stdin if empty)")
	flag.Parse()

	var r io.Reader
	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		r = f
	} else {
		r = os.Stdin
	}

	b, err := io.ReadAll(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading: %v\n", err)
		os.Exit(1)
	}

	_ = bufio.NewScanner(os.Stdin)
	text := string(b)

	printUniq("IPv4 Addresses", reIPv4.FindAllString(text, -1))
	printUniq("Domains", reDomain.FindAllString(text, -1))
	printUniq("URLs", reURL.FindAllString(text, -1))
	printUniq("Email Addresses", reEmail.FindAllString(text, -1))
	printUniq("MD5 Hashes", reMD5.FindAllString(text, -1))
	printUniq("SHA1 Hashes", reSHA1.FindAllString(text, -1))
	printUniq("SHA256 Hashes", reSHA256.FindAllString(text, -1))
}

func printUniq(label string, items []string) {
	if len(items) == 0 {
		return
	}
	seen := map[string]bool{}
	var uniq []string
	for _, v := range items {
		if !seen[v] {
			seen[v] = true
			uniq = append(uniq, v)
		}
	}
	sort.Strings(uniq)
	fmt.Printf("\n[%s] (%d)\n", label, len(uniq))
	for _, v := range uniq {
		fmt.Printf("  %s\n", v)
	}
}
