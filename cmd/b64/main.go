package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	decode := flag.Bool("d", false, "Decode instead of encode")
	input := flag.String("input", "", "Input string (reads from stdin if empty)")
	urlSafe := flag.Bool("url", false, "Use URL-safe encoding")
	flag.Parse()

	var raw string
	if *input != "" {
		raw = *input
	} else {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading stdin: %v\n", err)
			os.Exit(1)
		}
		raw = strings.TrimRight(string(b), "\n")
	}

	var enc *base64.Encoding
	if *urlSafe {
		enc = base64.URLEncoding
	} else {
		enc = base64.StdEncoding
	}

	if *decode {
		out, err := enc.DecodeString(raw)
		if err != nil {
			fmt.Fprintf(os.Stderr, "decode error: %v\n", err)
			os.Exit(1)
		}
		fmt.Print(string(out))
	} else {
		fmt.Println(enc.EncodeToString([]byte(raw)))
	}
}
