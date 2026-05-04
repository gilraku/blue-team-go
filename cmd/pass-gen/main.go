package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	lower   = "abcdefghijklmnopqrstuvwxyz"
	upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
	special = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func main() {
	length := flag.Int("n", 20, "Password length")
	count := flag.Int("count", 1, "Number of passwords to generate")
	noSpecial := flag.Bool("no-special", false, "Exclude special characters")
	noDigits := flag.Bool("no-digits", false, "Exclude digits")
	noUpper := flag.Bool("no-upper", false, "Exclude uppercase")
	flag.Parse()

	charset := lower
	if !*noUpper {
		charset += upper
	}
	if !*noDigits {
		charset += digits
	}
	if !*noSpecial {
		charset += special
	}

	if len(charset) == 0 {
		fmt.Fprintln(os.Stderr, "error: no character classes selected")
		os.Exit(1)
	}

	for i := 0; i < *count; i++ {
		pw, err := generate(charset, *length)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(pw)
	}

	if *count == 1 {
		fmt.Fprintf(os.Stderr, "\nCharset size: %d  |  Length: %d  |  Combinations: ~2^%.0f\n",
			len(charset), *length,
			float64(*length)*log2(float64(len(charset))))
	}
}

func generate(charset string, length int) (string, error) {
	b := make([]byte, length)
	max := big.NewInt(int64(len(charset)))
	for i := range b {
		idx, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		b[i] = charset[idx.Int64()]
	}
	_ = strings.Builder{}
	return string(b), nil
}

func log2(x float64) float64 {
	if x <= 0 {
		return 0
	}
	n := 0.0
	for x > 1 {
		x /= 2
		n++
	}
	return n
}
