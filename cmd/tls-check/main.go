package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	host := flag.String("host", "", "Host to check (e.g. example.com or example.com:443)")
	flag.Parse()

	if *host == "" {
		fmt.Fprintln(os.Stderr, "Usage: tls-check -host <hostname[:port]>")
		os.Exit(1)
	}

	addr := *host
	if !strings.Contains(addr, ":") {
		addr += ":443"
	}

	conn, err := tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: false})
	if err != nil {
		fmt.Fprintf(os.Stderr, "TLS error: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	state := conn.ConnectionState()
	certs := state.PeerCertificates

	fmt.Printf("Host:             %s\n", *host)
	fmt.Printf("TLS Version:      %s\n", tlsVersion(state.Version))
	fmt.Printf("Cipher Suite:     %s\n", tls.CipherSuiteName(state.CipherSuite))
	fmt.Printf("Certificates:     %d\n\n", len(certs))

	for i, cert := range certs {
		daysLeft := int(time.Until(cert.NotAfter).Hours() / 24)
		status := "VALID"
		if time.Now().After(cert.NotAfter) {
			status = "EXPIRED"
		} else if daysLeft < 30 {
			status = fmt.Sprintf("EXPIRING SOON (%d days)", daysLeft)
		}

		fmt.Printf("  Cert #%d:\n", i+1)
		fmt.Printf("    Subject:    %s\n", cert.Subject.CommonName)
		fmt.Printf("    Issuer:     %s\n", cert.Issuer.CommonName)
		fmt.Printf("    Valid From: %s\n", cert.NotBefore.Format("2006-01-02"))
		fmt.Printf("    Valid To:   %s\n", cert.NotAfter.Format("2006-01-02"))
		fmt.Printf("    Status:     %s\n", status)
		if len(cert.DNSNames) > 0 {
			fmt.Printf("    SANs:       %s\n", strings.Join(cert.DNSNames, ", "))
		}
		fmt.Println()
	}
}

func tlsVersion(v uint16) string {
	switch v {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return fmt.Sprintf("unknown (0x%04x)", v)
	}
}
