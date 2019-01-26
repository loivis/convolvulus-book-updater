package zongheng

import "testing"

func TestNew(t *testing.T) {
	t.Run("DefaultConfig", func(t *testing.T) {
		s := New()

		if s == nil {
			t.Fatal("s is nil")
		}
	})

	t.Run("WithName", func(t *testing.T) {
		str := "foo"
		s := New(
			WithName(str),
		)

		if got, want := s.name, str; got != want {
			t.Fatalf("s.name = %q, want %q", got, want)
		}
	})
}
