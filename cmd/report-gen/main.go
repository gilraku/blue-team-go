package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Finding struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Severity string `json:"severity"`
	Host     string `json:"host"`
	Port     int    `json:"port,omitempty"`
	Detail   string `json:"detail"`
	Remediation string `json:"remediation"`
}

type Report struct {
	Title    string    `json:"title"`
	Target   string    `json:"target"`
	Date     string    `json:"date"`
	Findings []Finding `json:"findings"`
}

func main() {
	input := flag.String("input", "", "JSON findings file")
	output := flag.String("output", "", "Output Markdown file (stdout if empty)")
	flag.Parse()

	if *input == "" {
		fmt.Fprintln(os.Stderr, "Usage: report-gen -input <findings.json> [-output <report.md>]")
		os.Exit(1)
	}

	data, err := os.ReadFile(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	var report Report
	if err := json.Unmarshal(data, &report); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	if report.Date == "" {
		report.Date = time.Now().Format("2006-01-02")
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# %s\n\n", report.Title))
	sb.WriteString(fmt.Sprintf("**Target:** %s  \n", report.Target))
	sb.WriteString(fmt.Sprintf("**Date:** %s  \n", report.Date))
	sb.WriteString(fmt.Sprintf("**Total Findings:** %d\n\n", len(report.Findings)))

	counts := map[string]int{"Critical": 0, "High": 0, "Medium": 0, "Low": 0, "Info": 0}
	for _, f := range report.Findings {
		counts[f.Severity]++
	}

	sb.WriteString("## Summary\n\n")
	sb.WriteString("| Severity | Count |\n|----------|-------|\n")
	for _, sev := range []string{"Critical", "High", "Medium", "Low", "Info"} {
		sb.WriteString(fmt.Sprintf("| %s | %d |\n", sev, counts[sev]))
	}

	sb.WriteString("\n## Findings\n\n")
	for _, f := range report.Findings {
		sb.WriteString(fmt.Sprintf("### [%s] %s\n\n", f.ID, f.Title))
		sb.WriteString(fmt.Sprintf("**Severity:** %s  \n", f.Severity))
		sb.WriteString(fmt.Sprintf("**Host:** %s", f.Host))
		if f.Port > 0 {
			sb.WriteString(fmt.Sprintf(":%d", f.Port))
		}
		sb.WriteString("\n\n")
		sb.WriteString(fmt.Sprintf("**Detail:**\n%s\n\n", f.Detail))
		sb.WriteString(fmt.Sprintf("**Remediation:**\n%s\n\n", f.Remediation))
		sb.WriteString("---\n\n")
	}

	md := sb.String()

	if *output != "" {
		if err := os.WriteFile(*output, []byte(md), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "error writing output: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Report written to %s\n", *output)
	} else {
		fmt.Print(md)
	}
}
