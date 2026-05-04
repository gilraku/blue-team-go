package main
import ("encoding/json";"testing")
func TestReportParse(t *testing.T) {
	raw := `{"title":"Test","target":"localhost","findings":[]}`
	var r Report
	if err := json.Unmarshal([]byte(raw), &r); err != nil { t.Fatal(err) }
	if r.Title != "Test" { t.Errorf("expected Test got %s", r.Title) }
}
