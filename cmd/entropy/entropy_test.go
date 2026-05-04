package main
import ("math";"testing")
func TestShannonEntropy(t *testing.T) {
	uniform := make([]byte, 256)
	for i := range uniform { uniform[i] = byte(i) }
	e := shannonEntropy(uniform)
	if math.Abs(e-8.0) > 0.01 { t.Errorf("uniform entropy should be ~8.0 got %.4f", e) }
	zeros := make([]byte, 100)
	e2 := shannonEntropy(zeros)
	if e2 != 0 { t.Errorf("all-zeros entropy should be 0 got %.4f", e2) }
}
