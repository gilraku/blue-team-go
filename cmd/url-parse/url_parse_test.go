package main
import ("net/url";"testing")
func TestURLParse(t *testing.T) {
	u, err := url.Parse("https://example.com/path?key=value")
	if err != nil { t.Fatal(err) }
	if u.Host != "example.com" { t.Errorf("expected example.com got %s", u.Host) }
	if u.Query().Get("key") != "value" { t.Errorf("expected value") }
}
