package algo

import (
	"testing"
)

func TestBruteForceBitParity(t *testing.T) {
	if BruteForceBitParity(254) != false {
		t.Fatal("expected odd bit parity")
	}
	if BruteForceBitParity(255) != true {
		t.Fatal("expected even bit parity")
	}

	if BruteForceBitParity(18446744073709551615) != true {
		t.Fatal("expected even bit parity")
	}

	if BruteForceBitParity(18446744073709551614) != false {
		t.Fatal("expected odd bit parity")
	}

}

func TestBitParity(t *testing.T) {
	if BitParity(254) != false {
		t.Fatal("expected odd bit parity")
	}
	if BitParity(255) != true {
		t.Fatal("expected even bit parity")
	}

	if BitParity(18446744073709551615) != true {
		t.Fatal("expected even bit parity")
	}

	if BitParity(18446744073709551614) != false {
		t.Fatal("expected odd bit parity")
	}

}
