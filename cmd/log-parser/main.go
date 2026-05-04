package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	reApache  = regexp.MustCompile(`^(\S+) \S+ \S+ \[([^\]]+)\] "(\S+) (\S+) \S+" (\d+) (\d+|-)`)
	reSyslog  = regexp.MustCompile(`^(\w{3}\s+\d+\s+\d+:\d+:\d+) (\S+) (\S+): (.+)$`)
	reSshFail = regexp.MustCompile(`Failed password for (?:invalid user )?(\S+) from (\S+)`)
)

type ApacheEntry struct {
	IP, Time, Method, Path, Status, Bytes string
}

func main() {
	file := flag.String("file", "", "Log file to parse")
	format := flag.String("format", "apache", "Log format: apache, syslog")
	filter := flag.String("filter", "", "Filter by status code or keyword")
	flag.Parse()

	if *file == "" {
		fmt.Fprintln(os.Stderr, "Usage: log-parser -file <path> [-format apache|syslog] [-filter <keyword>]")
		os.Exit(1)
	}

	f, err := os.Open(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var total, matched int

	for scanner.Scan() {
		line := scanner.Text()
		total++

		if *filter != "" && !strings.Contains(line, *filter) {
			continue
		}
		matched++

		switch *format {
		case "apache":
			if m := reApache.FindStringSubmatch(line); m != nil {
				fmt.Printf("[%s] %s %s %s → %s (%s bytes)\n", m[2], m[1], m[3], m[4], m[5], m[6])
			} else {
				fmt.Println(line)
			}
		case "syslog":
			if m := reSyslog.FindStringSubmatch(line); m != nil {
				fmt.Printf("[%s] %s %s: %s\n", m[1], m[2], m[3], m[4])
				if sf := reSshFail.FindStringSubmatch(m[4]); sf != nil {
					fmt.Printf("  !! SSH BRUTE FORCE: user=%s src=%s\n", sf[1], sf[2])
				}
			} else {
				fmt.Println(line)
			}
		default:
			fmt.Println(line)
		}
	}

	fmt.Printf("\n--- %d/%d lines matched ---\n", matched, total)
}
