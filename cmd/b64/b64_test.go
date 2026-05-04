package main
import ("encoding/base64";"testing")
func TestEncode(t *testing.T) {
	got := base64.StdEncoding.EncodeToString([]byte("hello world"))
	want := "aGVsbG8gd29ybGQ="
	if got != want { t.Errorf("got %s want %s", got, want) }
}
func TestDecode(t *testing.T) {
	b, _ := base64.StdEncoding.DecodeString("aGVsbG8gd29ybGQ=")
	if string(b) != "hello world" { t.Errorf("decode failed") }
}
