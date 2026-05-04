package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := flag.String("input", "", "String to encode (reads stdin if empty)")
	shift := flag.Int("n", 13, "Rotation shift (default 13 = ROT13)")
	all := flag.Bool("all", false, "Try all 25 rotations (brute force)")
	flag.Parse()

	var text string
	if *input != "" {
		text = *input
	} else {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			os.Exit(1)
		}
		text = strings.TrimRight(string(b), "\n")
	}

	if *all {
		for i := 1; i <= 25; i++ {
			fmt.Printf("ROT%-2d: %s\n", i, rotate(text, i))
		}
		return
	}

	fmt.Println(rotate(text, *shift))
}

func rotate(s string, n int) string {
	n = ((n % 26) + 26) % 26
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune('a' + (r-'a'+rune(n))%26)
		case r >= 'A' && r <= 'Z':
			b.WriteRune('A' + (r-'A'+rune(n))%26)
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}
