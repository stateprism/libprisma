package cryptoutil_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/stateprism/libprisma/cryptoutil"
	"testing"
)

func TestSeededRandomData(t *testing.T) {
	seed := []byte("seed")
	n := 10
	out := cryptoutil.SeededRandomData(seed, n)
	if len(out) != n {
		t.Errorf("Expected output length to be %d, got %d", n, len(out))
	}
	out2 := cryptoutil.SeededRandomData(seed, n)
	if !cmp.Equal(out, out2) {
		t.Errorf("Expected output to be the same for the same seed")
	}
}
