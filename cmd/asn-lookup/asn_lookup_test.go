package main
import ("net";"testing")
func TestParseIP(t *testing.T) {
	if net.ParseIP("8.8.8.8") == nil { t.Error("valid IP failed to parse") }
	if net.ParseIP("invalid") != nil { t.Error("invalid IP should be nil") }
}
