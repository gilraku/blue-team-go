package main
import "testing"
func TestRotate(t *testing.T) {
	if rotate("Hello", 13) != "Uryyb" { t.Errorf("ROT13 failed") }
	if rotate(rotate("Hello", 13), 13) != "Hello" { t.Error("double ROT13 should return original") }
	if rotate("abc", 0) != "abc" { t.Error("ROT0 should not change") }
}
