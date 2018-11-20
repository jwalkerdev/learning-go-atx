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

	areaTests := []struct {
		desc    string
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{12.0, 6.0}, hasArea: 72.0, desc: "Rectangles - 1"},
		{shape: Rectangle{0.0, 6.0}, hasArea: 0.0, desc: "Rectangles - Zero arg"},
		{shape: Circle{10.0}, hasArea: 314.1592653589793, desc: "Circles - 1"},
		{shape: Circle{0.0}, hasArea: 0.0, desc: "Circles - Zero arg"},
		{shape: Triangle{12, 6}, hasArea: 36.0, desc: "Triangles - 1"},
	}

	for _, test := range areaTests {
		t.Run(test.desc, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf("%#v got %.2f want %.2f", test.shape, got, test.hasArea)
			}
		})
	}
}
