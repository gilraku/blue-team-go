package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	file := flag.String("file", "", "File to hex dump (reads stdin if empty)")
	limit := flag.Int("n", 0, "Max bytes to dump (0 = all)")
	width := flag.Int("w", 16, "Bytes per row")
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

	buf := make([]byte, *width)
	offset := 0

	for {
		if *limit > 0 && offset >= *limit {
			break
		}

		toRead := *width
		if *limit > 0 && offset+toRead > *limit {
			toRead = *limit - offset
		}

		n, err := io.ReadFull(r, buf[:toRead])
		if n == 0 {
			break
		}

		row := buf[:n]
		fmt.Printf("%08X  ", offset)

		for i, b := range row {
			fmt.Printf("%02X ", b)
			if i == *width/2-1 {
				fmt.Print(" ")
			}
		}
		for i := n; i < *width; i++ {
			fmt.Print("   ")
			if i == *width/2-1 {
				fmt.Print(" ")
			}
		}

		fmt.Print(" |")
		for _, b := range row {
			if unicode.IsPrint(rune(b)) && b < 128 {
				fmt.Printf("%c", b)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")

		offset += n
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
	}

	fmt.Printf("%08X\n", offset)
}
