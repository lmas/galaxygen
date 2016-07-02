package galaxygen

import "testing"

var starTable = []struct {
	index int
	x     int
	y     int
}{
	// I just picked the 5 first/last stars for now...
	{0, -4, 22},
	{1, 0, 1},
	{2, -11, 9},
	{3, 0, 7},
	{4, 21, 10},
	{2304, -27, 35},
	{2305, 26, -32},
	{2306, -5, 65},
	{2307, -14, 26},
	{2308, 5, 9},
}

func TestGenerateGalaxy(t *testing.T) {
	galaxy := New(5000, 5, 100.0)

	if len(galaxy) < 1 {
		t.Error("Empty galaxy not expected")
	}

	for _, test := range starTable {
		star := galaxy[test.index]
		if star.X != test.x {
			t.Errorf("Expected X: %d, Got: %d\n", test.x, star.X)
		}
		if star.Y != test.y {
			t.Errorf("Expected Y: %d, Got: %d\n", test.y, star.Y)
		}
	}
}
