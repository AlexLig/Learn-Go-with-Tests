package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("y", 5)
	expected := "yyyyy"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
