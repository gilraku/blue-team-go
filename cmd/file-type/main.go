package main

import (
	"flag"
	"fmt"
	"os"
)

type signature struct {
	magic  []byte
	offset int
	desc   string
}

var signatures = []signature{
	{[]byte{0x4D, 0x5A}, 0, "PE/COFF Executable (Windows .exe/.dll)"},
	{[]byte{0x7F, 0x45, 0x4C, 0x46}, 0, "ELF Executable (Linux binary)"},
	{[]byte{0xFE, 0xED, 0xFA, 0xCE}, 0, "Mach-O 32-bit executable"},
	{[]byte{0xFE, 0xED, 0xFA, 0xCF}, 0, "Mach-O 64-bit executable"},
	{[]byte{0x25, 0x50, 0x44, 0x46}, 0, "PDF Document"},
	{[]byte{0x50, 0x4B, 0x03, 0x04}, 0, "ZIP Archive (or DOCX/XLSX/JAR)"},
	{[]byte{0x50, 0x4B, 0x05, 0x06}, 0, "ZIP Archive (empty)"},
	{[]byte{0x52, 0x61, 0x72, 0x21}, 0, "RAR Archive"},
	{[]byte{0x1F, 0x8B}, 0, "GZIP Compressed"},
	{[]byte{0x42, 0x5A, 0x68}, 0, "BZIP2 Compressed"},
	{[]byte{0xFD, 0x37, 0x7A, 0x58, 0x5A, 0x00}, 0, "XZ Compressed"},
	{[]byte{0xFF, 0xD8, 0xFF}, 0, "JPEG Image"},
	{[]byte{0x89, 0x50, 0x4E, 0x47}, 0, "PNG Image"},
	{[]byte{0x47, 0x49, 0x46, 0x38}, 0, "GIF Image"},
	{[]byte{0x4F, 0x67, 0x67, 0x53}, 0, "OGG Media"},
	{[]byte{0x49, 0x44, 0x33}, 0, "MP3 Audio (ID3)"},
	{[]byte{0x66, 0x74, 0x79, 0x70}, 4, "MP4 Video"},
	{[]byte{0xD0, 0xCF, 0x11, 0xE0}, 0, "MS Office (DOC/XLS/PPT - OLE2)"},
	{[]byte{0x37, 0x7A, 0xBC, 0xAF, 0x27, 0x1C}, 0, "7-Zip Archive"},
	{[]byte{0x75, 0x73, 0x74, 0x61, 0x72}, 257, "TAR Archive"},
	{[]byte{0x4D, 0x53, 0x43, 0x46}, 0, "Microsoft Cabinet (.cab)"},
	{[]byte{0x23, 0x21}, 0, "Script (shebang)"},
	{[]byte{0xCA, 0xFE, 0xBA, 0xBE}, 0, "Java Class File"},
}

func main() {
	file := flag.String("file", "", "File to identify")
	flag.Parse()

	if *file == "" {
		fmt.Fprintln(os.Stderr, "Usage: file-type -file <path>")
		os.Exit(1)
	}

	f, err := os.Open(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil && n == 0 {
		fmt.Fprintf(os.Stderr, "error reading: %v\n", err)
		os.Exit(1)
	}
	buf = buf[:n]

	info, _ := f.Stat()
	fmt.Printf("File:   %s\n", *file)
	if info != nil {
		fmt.Printf("Size:   %d bytes\n", info.Size())
	}

	for _, sig := range signatures {
		if sig.offset+len(sig.magic) > len(buf) {
			continue
		}
		match := true
		for i, b := range sig.magic {
			if buf[sig.offset+i] != b {
				match = false
				break
			}
		}
		if match {
			fmt.Printf("Type:   %s\n", sig.desc)
			return
		}
	}

	printable := true
	for _, b := range buf[:min(buf, 64)] {
		if b < 0x09 || (b > 0x0D && b < 0x20) {
			printable = false
			break
		}
	}
	if printable {
		fmt.Println("Type:   Text/ASCII (no magic bytes matched)")
	} else {
		fmt.Println("Type:   Unknown binary")
	}
}

func min(data []byte, n int) int {
	if len(data) < n {
		return len(data)
	}
	return n
}
