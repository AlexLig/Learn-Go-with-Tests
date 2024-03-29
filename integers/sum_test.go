package integers

import "testing"

func TestSum(t *testing.T) {
	t.Run("number collection of any size", func(t *testing.T) {
		numbers := []int{3, 2, 1}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
