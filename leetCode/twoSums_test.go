package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSums(t *testing.T) {
	t.Run("return indices of the two numbers such that they add up to target", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		target := 3
		got := twoSum(input, target)
		want := []int{0, 1}
		assertIndices(t, got, want)
	})

}

func assertIndices(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
