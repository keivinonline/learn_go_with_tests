package main

import "testing"

func assertCorrectResult(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want: %v", got, want)
	}
}

func TestPerimenter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(&rectangle)
	want := 40.0
	assertCorrectResult(t, got, want)
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("Area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 12.0, Height: 6.0}
		checkArea(t, rectangle, 72.0)
	})
	t.Run("Area of a circle", func(t *testing.T) {
		circle := Circle{Radius: 10}
		checkArea(t, circle, 314.1592653589793)
	})
	//	func TestArea(t *testing.T) {
	//		t.Run("Area of a rectangle", func(t *testing.T) {
	//			rectangle := Rectangle{Width: 12.0, Height: 6.0}
	//			got := rectangle.Area()
	//			want := 72.0
	//			assertCorrectResult(t, got, want)
	//		})
	//		t.Run("Area of a circle", func(t *testing.T) {
	//			circle := Circle{Radius: 10}
	//			got := circle.Area()
	//			want := 314.1592653589793
	//			assertCorrectResult(t, got, want)
	//		})
	//	}
}

func TestArea2(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v - got %g want %g",tt.shape, got, tt.want)
		}
	}

}
