package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type Connection struct {
	Proto, LocalAddr, ForeignAddr, State string
}

func main() {
	file := flag.String("file", "", "netstat output file (reads stdin if empty)")
	state := flag.String("state", "", "Filter by state (LISTEN, ESTABLISHED, TIME_WAIT, etc.)")
	summary := flag.Bool("summary", false, "Show state summary only")
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

	scanner := bufio.NewScanner(r)
	var conns []Connection
	stateCounts := map[string]int{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}

		proto := strings.ToLower(fields[0])
		if proto != "tcp" && proto != "tcp6" && proto != "udp" && proto != "udp6" {
			continue
		}

		c := Connection{Proto: fields[0]}
		if len(fields) >= 4 {
			c.LocalAddr = fields[3]
		}
		if len(fields) >= 5 {
			c.ForeignAddr = fields[4]
		}
		if len(fields) >= 6 {
			c.State = fields[5]
		}

		stateCounts[c.State]++

		if *state != "" && !strings.EqualFold(c.State, *state) {
			continue
		}
		conns = append(conns, c)
	}

	if *summary {
		states := make([]string, 0, len(stateCounts))
		for s := range stateCounts {
			states = append(states, s)
		}
		sort.Strings(states)
		fmt.Printf("%-20s %s\n", "State", "Count")
		fmt.Printf("%-20s %s\n", "-----", "-----")
		for _, s := range states {
			fmt.Printf("%-20s %d\n", s, stateCounts[s])
		}
		return
	}

	fmt.Printf("%-8s %-30s %-30s %s\n", "Proto", "Local Address", "Foreign Address", "State")
	fmt.Printf("%-8s %-30s %-30s %s\n", "-----", "-------------", "---------------", "-----")
	for _, c := range conns {
		fmt.Printf("%-8s %-30s %-30s %s\n", c.Proto, c.LocalAddr, c.ForeignAddr, c.State)
	}
	fmt.Printf("\nTotal: %d connections\n", len(conns))
}
