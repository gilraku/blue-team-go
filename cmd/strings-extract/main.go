package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	file := flag.String("file", "", "File to extract strings from (reads stdin if empty)")
	minLen := flag.Int("n", 4, "Minimum string length")
	printOffset := flag.Bool("o", false, "Print file offset")
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

	data, err := io.ReadAll(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	var current []byte
	start := 0

	for i, b := range data {
		if unicode.IsPrint(rune(b)) && b < 128 {
			if len(current) == 0 {
				start = i
			}
			current = append(current, b)
		} else {
			if len(current) >= *minLen {
				if *printOffset {
					fmt.Printf("%08X  %s\n", start, string(current))
				} else {
					fmt.Println(string(current))
				}
			}
			current = current[:0]
		}
	}

	if len(current) >= *minLen {
		if *printOffset {
			fmt.Printf("%08X  %s\n", start, string(current))
		} else {
			fmt.Println(string(current))
		}
	}
}
