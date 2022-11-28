package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	expect := 15
	if got != expect {
		t.Errorf("got %d expect %d given %v", got, expect, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6, 7})
	expect := []int{6, 22}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got %v expect %v", got, expect)
	}
}

func TestSumAllTails(t *testing.T) {

	assertSliceEqual := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("safely sum slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{5, 6})
		want := []int{5, 6}
		assertSliceEqual(t, got, want)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1}, []int{1, 2}, []int{1, 2, 3})
		want := []int{0, 0, 2, 5}
		assertSliceEqual(t, got, want)
	})
}
