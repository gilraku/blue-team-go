package main
import "testing"
func TestUAParse(t *testing.T) {
	info := parse("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	if info.Browser != "Google Chrome" { t.Errorf("expected Chrome got %s", info.Browser) }
	if info.OS != "Windows" { t.Errorf("expected Windows got %s", info.OS) }
}
func TestBotDetect(t *testing.T) {
	info := parse("Googlebot/2.1 (+http://www.google.com/bot.html)")
	if info.DeviceType != "Bot/Crawler" { t.Errorf("expected Bot got %s", info.DeviceType) }
}
