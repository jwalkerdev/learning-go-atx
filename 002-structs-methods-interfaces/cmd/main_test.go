package main

import (
	"testing"

	"github.com/jwalkerdev/learning-go-class/002-structs-methods-interfaces/shape"
)

func TestGetArea(t *testing.T) {
	t.Run("Rectangle Area", func(t *testing.T) {
		got := getArea(shape.Rectangle{10, 20})
		want := 200.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
