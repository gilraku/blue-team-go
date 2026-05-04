package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type UAInfo struct {
	Browser, BrowserVersion, OS, OSVersion, DeviceType, Engine string
}

func parse(ua string) UAInfo {
	info := UAInfo{DeviceType: "Desktop"}

	if strings.Contains(ua, "Mobile") || strings.Contains(ua, "Android") {
		info.DeviceType = "Mobile"
	}
	if strings.Contains(ua, "Tablet") || strings.Contains(ua, "iPad") {
		info.DeviceType = "Tablet"
	}
	if strings.Contains(ua, "bot") || strings.Contains(ua, "Bot") || strings.Contains(ua, "crawler") {
		info.DeviceType = "Bot/Crawler"
	}

	switch {
	case strings.Contains(ua, "Edg/") || strings.Contains(ua, "Edge/"):
		info.Browser = "Microsoft Edge"
		info.Engine = "Blink"
		info.BrowserVersion = extract(ua, "Edg/", " ")
	case strings.Contains(ua, "OPR/") || strings.Contains(ua, "Opera/"):
		info.Browser = "Opera"
		info.Engine = "Blink"
		info.BrowserVersion = extract(ua, "OPR/", " ")
	case strings.Contains(ua, "Chrome/") && !strings.Contains(ua, "Chromium"):
		info.Browser = "Google Chrome"
		info.Engine = "Blink"
		info.BrowserVersion = extract(ua, "Chrome/", " ")
	case strings.Contains(ua, "Chromium/"):
		info.Browser = "Chromium"
		info.Engine = "Blink"
		info.BrowserVersion = extract(ua, "Chromium/", " ")
	case strings.Contains(ua, "Firefox/"):
		info.Browser = "Mozilla Firefox"
		info.Engine = "Gecko"
		info.BrowserVersion = extract(ua, "Firefox/", " ")
	case strings.Contains(ua, "Safari/") && !strings.Contains(ua, "Chrome"):
		info.Browser = "Safari"
		info.Engine = "WebKit"
		info.BrowserVersion = extract(ua, "Version/", " ")
	case strings.Contains(ua, "curl/"):
		info.Browser = "curl"
		info.BrowserVersion = extract(ua, "curl/", " ")
	case strings.Contains(ua, "python-requests"):
		info.Browser = "Python Requests"
	case strings.Contains(ua, "Wget/"):
		info.Browser = "Wget"
	default:
		info.Browser = "Unknown"
	}

	switch {
	case strings.Contains(ua, "Windows NT 10.0"):
		info.OS = "Windows"
		info.OSVersion = "10/11"
	case strings.Contains(ua, "Windows NT 6.3"):
		info.OS = "Windows"
		info.OSVersion = "8.1"
	case strings.Contains(ua, "Windows NT 6.1"):
		info.OS = "Windows"
		info.OSVersion = "7"
	case strings.Contains(ua, "Mac OS X"):
		info.OS = "macOS"
		info.OSVersion = extract(ua, "Mac OS X ", ")")
	case strings.Contains(ua, "Android"):
		info.OS = "Android"
		info.OSVersion = extract(ua, "Android ", ";")
	case strings.Contains(ua, "iPhone OS") || strings.Contains(ua, "iPad; CPU OS"):
		info.OS = "iOS"
		info.OSVersion = strings.ReplaceAll(extract(ua, "OS ", " like"), "_", ".")
	case strings.Contains(ua, "Linux"):
		info.OS = "Linux"
	default:
		info.OS = "Unknown"
	}

	return info
}

func extract(s, after, before string) string {
	i := strings.Index(s, after)
	if i < 0 {
		return ""
	}
	s = s[i+len(after):]
	if j := strings.Index(s, before); j >= 0 {
		s = s[:j]
	}
	return strings.TrimSpace(s)
}

func main() {
	ua := flag.String("ua", "", "User-Agent string to parse")
	flag.Parse()

	if *ua == "" {
		fmt.Fprintln(os.Stderr, "Usage: ua-parse -ua <user-agent>")
		os.Exit(1)
	}

	info := parse(*ua)
	fmt.Printf("Browser:  %s %s\n", info.Browser, info.BrowserVersion)
	fmt.Printf("Engine:   %s\n", info.Engine)
	fmt.Printf("OS:       %s %s\n", info.OS, info.OSVersion)
	fmt.Printf("Device:   %s\n", info.DeviceType)
	fmt.Printf("\nRaw:      %s\n", *ua)
}
