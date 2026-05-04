package main
import "testing"
func TestMinHelper(t *testing.T) {
	data := []byte{1,2,3}
	if min(data, 5) != 3 { t.Error("min should return len(data) when n > len") }
	if min(data, 2) != 2 { t.Error("min should return n when n < len") }
}
