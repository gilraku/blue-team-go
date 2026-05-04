package main
import "testing"
func TestDetectServer(t *testing.T) {
	if detectServer("192.168.1.1") != "whois.arin.net" { t.Error("IP should use arin") }
	if detectServer("example.com") != "whois.verisign-grs.com" { t.Error(".com should use verisign") }
}
