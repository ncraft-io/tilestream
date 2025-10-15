package tilestream

import "testing"

func TestPolygonArea(t *testing.T) {
	xs := []float64{1, 2, 1}
	ys := []float64{1, 1, 2}

	t.Logf("area %f", _area(xs, ys, 3)) // -0.5

	xs = []float64{1, 2, 3, 2, 1}
	ys = []float64{1, 1, 2, 3, 2}

	t.Logf("area %f", _area(xs, ys, 5)) // -2.5

	xs = []float64{2, 3, 2, 1, 2}
	ys = []float64{1, 2, 3, 2, 2}

	t.Logf("area %f", _area(xs, ys, 5)) // -1.5

	xs = []float64{2, 1, 2, 3, 2}
	ys = []float64{2, 2, 3, 2, 1}

	t.Logf("area %f", _area(xs, ys, 5)) // 1.5
}
