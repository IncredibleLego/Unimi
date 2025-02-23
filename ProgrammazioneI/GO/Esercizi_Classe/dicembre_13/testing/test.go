//Autore: Francesco Corrado

package point

import(
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1E-6
}

func TestMedian(t *testing.T) {
	var p1, p2, pm point
	p1 = NewPoint(0, 0)
	p2 = NewPoint(2, 4)
	pm = Median(p1, p2)
	if !(almostEqual(pm.X, 1) && almostEqual(pm.Y, 2)){
		t.Error("Expected (1, 2) got ", pm)
	}
}