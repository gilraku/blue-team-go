package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	token := flag.String("token", "", "JWT token to decode")
	flag.Parse()

	if *token == "" {
		fmt.Fprintln(os.Stderr, "Usage: jwt-decode -token <jwt>")
		os.Exit(1)
	}

	parts := strings.Split(*token, ".")
	if len(parts) != 3 {
		fmt.Fprintln(os.Stderr, "invalid JWT: must have 3 parts")
		os.Exit(1)
	}

	header, err := decodeSegment(parts[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error decoding header: %v\n", err)
		os.Exit(1)
	}

	payload, err := decodeSegment(parts[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error decoding payload: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("=== HEADER ===")
	printJSON(header)

	fmt.Println("\n=== PAYLOAD ===")
	printJSON(payload)

	var claims map[string]any
	if err := json.Unmarshal(payload, &claims); err == nil {
		fmt.Println("\n=== TIMESTAMPS ===")
		for _, field := range []string{"iat", "exp", "nbf"} {
			if v, ok := claims[field]; ok {
				if ts, ok := v.(float64); ok {
					t := time.Unix(int64(ts), 0).UTC()
					label := map[string]string{
						"iat": "Issued At ",
						"exp": "Expires At",
						"nbf": "Not Before",
					}[field]
					status := ""
					if field == "exp" {
						if t.Before(time.Now()) {
							status = " [EXPIRED]"
						} else {
							status = " [VALID]"
						}
					}
					fmt.Printf("  %s: %s%s\n", label, t.Format(time.RFC3339), status)
				}
			}
		}
	}

	fmt.Println("\n=== SIGNATURE ===")
	fmt.Println("  (not verified — use a proper library for signature validation)")
}

func decodeSegment(s string) ([]byte, error) {
	padded := s + strings.Repeat("=", (4-len(s)%4)%4)
	return base64.URLEncoding.DecodeString(padded)
}

func printJSON(data []byte) {
	var out map[string]any
	if err := json.Unmarshal(data, &out); err != nil {
		fmt.Println(string(data))
		return
	}
	b, _ := json.MarshalIndent(out, "  ", "  ")
	fmt.Println(" ", string(b))
}
