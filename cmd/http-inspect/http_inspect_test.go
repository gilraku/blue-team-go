package main
import "testing"
func TestSecurityHeaders(t *testing.T) {
	if len(securityHeaders) == 0 { t.Error("security headers list must not be empty") }
}
