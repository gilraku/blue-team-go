package main
import "testing"
func TestIPRegex(t *testing.T) {
	matches := reIPv4.FindAllString("Alert from 192.168.1.1 and 10.0.0.1", -1)
	if len(matches) != 2 { t.Errorf("expected 2 IPs got %d", len(matches)) }
}
func TestDomainRegex(t *testing.T) {
	matches := reDomain.FindAllString("Visit evil.com and malware.io now", -1)
	if len(matches) == 0 { t.Error("expected domains") }
}
