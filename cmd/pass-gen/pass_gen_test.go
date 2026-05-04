package main
import "testing"
func TestGenerate(t *testing.T) {
	pw, err := generate("abcdefghij", 16)
	if err != nil { t.Fatal(err) }
	if len(pw) != 16 { t.Errorf("expected length 16 got %d", len(pw)) }
	for _, c := range pw {
		if c < a || c > j { t.Errorf("unexpected char %c", c) }
	}
}
