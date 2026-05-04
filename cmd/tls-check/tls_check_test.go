package main
import ("crypto/tls";"testing")
func TestTLSVersion(t *testing.T) {
	if tlsVersion(tls.VersionTLS13) != "TLS 1.3" { t.Error("wrong TLS 1.3 label") }
	if tlsVersion(tls.VersionTLS12) != "TLS 1.2" { t.Error("wrong TLS 1.2 label") }
}
