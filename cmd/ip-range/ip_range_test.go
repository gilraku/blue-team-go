package main
import ("net";"testing")
func TestCloneIP(t *testing.T) {
	orig := net.IP{10,0,0,1}
	clone := cloneIP(orig)
	clone[3] = 99
	if orig[3] != 1 { t.Error("cloneIP modified original") }
}
func TestInc(t *testing.T) {
	ip := net.IP{10,0,0,1}
	inc(ip)
	if ip[3] != 2 { t.Errorf("expected 2 got %d", ip[3]) }
}
