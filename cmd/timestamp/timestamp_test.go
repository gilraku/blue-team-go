package main
import ("testing";"time")
func TestUnixConvert(t *testing.T) {
	ts := int64(0)
	got := time.Unix(ts, 0).UTC().Format("2006-01-02")
	if got != "1970-01-01" { t.Errorf("expected 1970-01-01 got %s", got) }
}
