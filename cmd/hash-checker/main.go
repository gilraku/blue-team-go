package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	algo := flag.String("algo", "sha256", "Hash algorithm: md5, sha1, sha256, sha512")
	input := flag.String("input", "", "String to hash")
	file := flag.String("file", "", "File to hash")
	compare := flag.String("compare", "", "Hash to compare against")
	flag.Parse()

	if *input == "" && *file == "" {
		fmt.Fprintln(os.Stderr, "Usage: hash-checker -algo <algo> [-input <string>|-file <path>] [-compare <hash>]")
		os.Exit(1)
	}

	var data []byte
	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		data, err = io.ReadAll(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
			os.Exit(1)
		}
	} else {
		data = []byte(*input)
	}

	var result string
	switch *algo {
	case "md5":
		h := md5.Sum(data)
		result = fmt.Sprintf("%x", h)
	case "sha1":
		h := sha1.Sum(data)
		result = fmt.Sprintf("%x", h)
	case "sha256":
		h := sha256.Sum256(data)
		result = fmt.Sprintf("%x", h)
	case "sha512":
		h := sha512.Sum512(data)
		result = fmt.Sprintf("%x", h)
	default:
		fmt.Fprintf(os.Stderr, "unknown algorithm: %s\n", *algo)
		os.Exit(1)
	}

	fmt.Printf("%s  %s\n", result, func() string {
		if *file != "" {
			return *file
		}
		return "-"
	}())

	if *compare != "" {
		if result == *compare {
			fmt.Println("MATCH")
		} else {
			fmt.Println("MISMATCH")
			os.Exit(1)
		}
	}
}
