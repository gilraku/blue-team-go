package main
import "testing"
func TestHasCommonPattern(t *testing.T) {
	if !hasCommonPattern("Password123") { t.Error("should detect common pattern") }
	if hasCommonPattern("xK9$mQ2!vLpN") { t.Error("should not flag strong password") }
}
