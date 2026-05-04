package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func main() {
	pw := flag.String("pw", "", "Password to analyze")
	flag.Parse()

	if *pw == "" {
		fmt.Fprintln(os.Stderr, "Usage: pwcheck -pw <password>")
		os.Exit(1)
	}

	p := *pw
	var hasLower, hasUpper, hasDigit, hasSpecial bool
	charsetSize := 0

	for _, c := range p {
		switch {
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsDigit(c):
			hasDigit = true
		default:
			hasSpecial = true
		}
	}

	if hasLower {
		charsetSize += 26
	}
	if hasUpper {
		charsetSize += 26
	}
	if hasDigit {
		charsetSize += 10
	}
	if hasSpecial {
		charsetSize += 32
	}

	entropy := float64(len(p)) * math.Log2(float64(charsetSize))

	strength := "Very Weak"
	switch {
	case entropy >= 128:
		strength = "Very Strong"
	case entropy >= 80:
		strength = "Strong"
	case entropy >= 60:
		strength = "Moderate"
	case entropy >= 40:
		strength = "Weak"
	}

	fmt.Printf("Length:      %d\n", len(p))
	fmt.Printf("Charset:     %d chars\n", charsetSize)
	fmt.Printf("Entropy:     %.1f bits\n", entropy)
	fmt.Printf("Strength:    %s\n\n", strength)

	checks := []struct {
		label string
		pass  bool
	}{
		{"Lowercase letters", hasLower},
		{"Uppercase letters", hasUpper},
		{"Digits", hasDigit},
		{"Special characters", hasSpecial},
		{"Length >= 12", len(p) >= 12},
		{"Length >= 16", len(p) >= 16},
		{"No common patterns", !hasCommonPattern(p)},
	}

	for _, c := range checks {
		mark := "✓"
		if !c.pass {
			mark = "✗"
		}
		fmt.Printf("  [%s] %s\n", mark, c.label)
	}
}

func hasCommonPattern(p string) bool {
	lower := strings.ToLower(p)
	patterns := []string{"password", "123456", "qwerty", "abc123", "letmein", "admin", "welcome"}
	for _, pat := range patterns {
		if strings.Contains(lower, pat) {
			return true
		}
	}
	return false
}
