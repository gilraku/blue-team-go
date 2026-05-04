package main
import ("net";"testing")
func TestNextIP(t *testing.T) {
	ip := net.IP{192,168,1,0}
	next := nextIP(ip)
	if next[3] != 1 { t.Errorf("expected 1 got %d", next[3]) }
}
func TestPrevIP(t *testing.T) {
	ip := net.IP{192,168,1,255}
	prev := prevIP(ip)
	if prev[3] != 254 { t.Errorf("expected 254 got %d", prev[3]) }
}
