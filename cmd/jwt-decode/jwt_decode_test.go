package main
import "testing"
func TestDecodeSegment(t *testing.T) {
	out, err := decodeSegment("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	if err != nil { t.Fatal(err) }
	if len(out) == 0 { t.Error("expected non-empty output") }
}
