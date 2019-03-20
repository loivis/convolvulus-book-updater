package update

import "testing"

// TODO: add more tests
func TestBook_Equals(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		b1 := &Book{Author: "foo"}
		b2 := &Book{Author: "foo"}
		b3 := &Book{Author: "bar"}

		if !b1.Equals(b2) {
			t.Fatalf("equal wanted: %v vs %v", b1, b2)
		}

		if b1.Equals(b3) {
			t.Fatalf("unqual wanted: %v vs %v", b1, b2)
		}
	})
}
