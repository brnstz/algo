package algo

import (
	"testing"
)

func TestBitParity(t *testing.T) {
	if BruteForceBitParity(254) != 7 {
		t.Fatal("expected bit parity of 7")
	}
	if BruteForceBitParity(255) != 8 {
		t.Fatal("expected bit parity of 8")
	}

	if BruteForceBitParity(18446744073709551615) != 64 {
		t.Fatal("expected bit parity of 64")
	}

	if BruteForceBitParity(18446744073709551614) != 63 {
		t.Fatal("expected bit parity of 64")
	}

}
