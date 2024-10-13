package iteration

import "testing"

func assertCorrectResult(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	want := "aaaaaa"
	assertCorrectResult(t, got, want)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}

}
