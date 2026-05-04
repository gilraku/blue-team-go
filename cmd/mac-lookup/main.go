package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var ouiDB = map[string]string{
	"00:00:0C": "Cisco Systems",
	"00:1A:A0": "Dell",
	"00:50:56": "VMware",
	"00:0C:29": "VMware (Workstation)",
	"00:05:69": "VMware",
	"08:00:27": "Oracle VirtualBox",
	"52:54:00": "QEMU/KVM",
	"00:16:3E": "Xen",
	"00:1B:21": "Intel",
	"00:1C:BF": "Intel",
	"00:21:6A": "Intel",
	"00:1E:67": "Intel",
	"3C:97:0E": "Apple",
	"A4:C3:F0": "Apple",
	"F0:18:98": "Apple",
	"AC:DE:48": "Apple",
	"00:17:F2": "Apple",
	"DC:A9:04": "Huawei",
	"00:E0:FC": "Huawei",
	"54:89:98": "Huawei",
	"00:1A:2B": "Cisco-Linksys",
	"00:18:4D": "Netgear",
	"20:E5:2A": "Netgear",
	"00:26:F2": "Netgear",
	"B0:7F:B9": "TP-Link",
	"50:C7:BF": "TP-Link",
	"EC:08:6B": "TP-Link",
	"00:90:F5": "Mikrotik",
	"CC:2D:E0": "Mikrotik",
	"00:0D:B9": "PC Engines",
	"00:1D:AA": "D-Link",
	"1C:7E:E5": "D-Link",
	"00:17:9A": "D-Link",
	"00:1E:C2": "Ubiquiti",
	"04:18:D6": "Ubiquiti",
	"44:D9:E7": "Ubiquiti",
	"00:15:5D": "Microsoft (Hyper-V)",
	"00:03:FF": "Microsoft",
}

func main() {
	mac := flag.String("mac", "", "MAC address to look up (e.g. 00:50:56:AA:BB:CC)")
	flag.Parse()

	if *mac == "" {
		fmt.Fprintln(os.Stderr, "Usage: mac-lookup -mac <address>")
		os.Exit(1)
	}

	normalized := strings.ToUpper(strings.ReplaceAll(*mac, "-", ":"))
	parts := strings.Split(normalized, ":")
	if len(parts) < 3 {
		fmt.Fprintln(os.Stderr, "invalid MAC address format")
		os.Exit(1)
	}

	oui := strings.Join(parts[:3], ":")
	fmt.Printf("MAC:     %s\n", normalized)
	fmt.Printf("OUI:     %s\n", oui)

	if vendor, ok := ouiDB[oui]; ok {
		fmt.Printf("Vendor:  %s\n", vendor)
	} else {
		fmt.Println("Vendor:  Unknown (not in local database)")
		fmt.Println("Hint:    Try https://macvendors.com for full lookup")
	}

	if parts[0][1] == '2' || parts[0][1] == '6' || parts[0][1] == 'A' || parts[0][1] == 'E' {
		fmt.Println("Note:    Locally administered / randomized MAC address")
	}
}
