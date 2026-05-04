package main
import ("crypto/sha256";"fmt";"testing")
func TestSHA256(t *testing.T) {
	data := []byte("hello")
	h := sha256.Sum256(data)
	got := fmt.Sprintf("%x", h)
	want := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	if got != want { t.Errorf("got %s want %s", got, want) }
}
