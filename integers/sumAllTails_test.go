package integers

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {

	assertTailsSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Sums the tails of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 9}, []int{0, 1})
		want := []int{9, 1}

		assertTailsSum(t, got, want)
	})

	t.Run("Safely sums the tails of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 1, 1, 3})
		want := []int{0, 5}

		assertTailsSum(t, got, want)
	})

}
