package main
import "testing"
func TestOUILookup(t *testing.T) {
	if vendor, ok := ouiDB["00:50:56"]; !ok || vendor != "VMware" {
		t.Errorf("expected VMware got %s", vendor)
	}
}
