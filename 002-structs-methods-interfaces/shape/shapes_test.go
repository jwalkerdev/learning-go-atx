package shape

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	areaTests := []struct {
		shape Shape
		want  float64
		desc  string
	}{
		{Rectangle{12.0, 6.0}, 72.0, "Rectangles - 1"},
		{Rectangle{6.0, 6.0}, 36.0, "Rectangles - 2"},
		{Rectangle{0.0, 6.0}, 0.0, "Rectangles - Zero arg"},
		{Circle{10.0}, 314.1592653589793, "Circles - 1"},
		{Circle{100.0}, 31415.9265358979319, "Circles - 2"},
		{Circle{0.0}, 0.0, "Circles - Zero arg"},
	}

	for _, test := range areaTests {
		t.Run(test.desc, func(t *testing.T) {
			checkArea(t, test.shape, test.want)
		})
	}
}
