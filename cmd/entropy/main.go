package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func main() {
	file := flag.String("file", "", "File to calculate entropy for")
	blockSize := flag.Int("block", 0, "Block size for per-block analysis (0 = whole file)")
	flag.Parse()

	if *file == "" {
		fmt.Fprintln(os.Stderr, "Usage: entropy -file <path> [-block <size>]")
		os.Exit(1)
	}

	data, err := os.ReadFile(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if *blockSize > 0 {
		fmt.Printf("%-12s %s\n", "Offset", "Entropy")
		fmt.Printf("%-12s %s\n", "------", "-------")
		for i := 0; i < len(data); i += *blockSize {
			end := i + *blockSize
			if end > len(data) {
				end = len(data)
			}
			e := shannonEntropy(data[i:end])
			bar := makeBar(e)
			fmt.Printf("0x%-10X %.4f  %s\n", i, e, bar)
		}
	} else {
		e := shannonEntropy(data)
		fmt.Printf("File:    %s\n", *file)
		fmt.Printf("Size:    %d bytes\n", len(data))
		fmt.Printf("Entropy: %.4f bits/byte\n", e)
		fmt.Printf("Bar:     %s\n", makeBar(e))
		fmt.Printf("Note:    %s\n", interpret(e))
	}
}

func shannonEntropy(data []byte) float64 {
	if len(data) == 0 {
		return 0
	}
	freq := make([]int, 256)
	for _, b := range data {
		freq[b]++
	}
	n := float64(len(data))
	var h float64
	for _, f := range freq {
		if f > 0 {
			p := float64(f) / n
			h -= p * math.Log2(p)
		}
	}
	return h
}

func makeBar(e float64) string {
	width := int(e / 8.0 * 40)
	bar := ""
	for i := 0; i < 40; i++ {
		if i < width {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return bar
}

func interpret(e float64) string {
	switch {
	case e > 7.5:
		return "Very high — likely encrypted, compressed, or packed"
	case e > 6.5:
		return "High — possibly compressed or obfuscated"
	case e > 5.0:
		return "Medium — mixed binary/text data"
	case e > 3.5:
		return "Low-medium — structured binary or rich text"
	default:
		return "Low — plain text or highly structured data"
	}
}
