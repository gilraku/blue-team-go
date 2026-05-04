package main
import "testing"
func TestApacheRegex(t *testing.T) {
	line := `192.168.1.1 - - [04/May/2026:10:00:00 +0000] "GET /index.html HTTP/1.1" 200 1234`
	m := reApache.FindStringSubmatch(line)
	if m == nil { t.Error("expected apache match") }
	if m[1] != "192.168.1.1" { t.Errorf("expected IP 192.168.1.1 got %s", m[1]) }
}
