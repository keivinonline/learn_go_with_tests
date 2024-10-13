package main

import (
	"slices"
	"testing"
)

func assertCorrectResult(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got :%v, want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{
			1, 2, 3, 4, 5,
		}
		got := Sum(numbers)
		want := 15
		assertCorrectResult(t, got, want)
	})
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{
			1, 2, 3,
		}
		got := Sum(numbers)
		want := 6
		assertCorrectResult(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sums of some slices", func(t *testing.T) {
		numbers1 := []int{1, 2}
		numbers2 := []int{0, 9}
		got := SumAllTails(numbers1, numbers2)
		want := []int{2, 9}

		// if reflect.DeepEqual(got, want) != true {
		// 	t.Errorf("got %v, ant %v", got, want)
		// }

		if slices.Compare(got, want) != 0 {
			t.Errorf("got %v, ant %v", got, want)
		}
	})
	t.Run("safetly sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		// if reflect.DeepEqual(got, want) != true {
		// 	t.Errorf("got %v, ant %v", got, want)
		// }

		if slices.Compare(got, want) != 0 {
			t.Errorf("got %v, ant %v", got, want)
		}
	})

}
