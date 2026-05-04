package main
import "testing"
func TestMin2(t *testing.T) {
	if min2(3,5) != 3 { t.Error("min2 wrong") }
	if min2(7,4) != 4 { t.Error("min2 wrong") }
}
