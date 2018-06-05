package algo

import (
	"log"
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

func TestBitSwap(t *testing.T) {
	var x int64

	x = BitSwap(0x20, 2, 3)
	if x != 0x20 {
		t.Fatalf("expected %b but got %b", 0x20, x)
	}

	x = BitSwap(0x1a, 3, 2)
	if x != 0x16 {
		t.Fatalf("expected %b but got %b", 0x16, x)
	}

}

func TestBitReverse(t *testing.T) {
	if BitReverse(1) != 9223372036854775808 {
		log.Fatal("bit reverse failed")
	}
}

func TestBitAdd(t *testing.T) {
	var x uint64

	x = BitAdd(1, 3)
	if x != 4 {
		t.Fatalf("expected 1 + 3 = 4 but got %v", x)
	}

	x = BitAdd(0, 0)
	if x != 0 {
		t.Fatalf("expected 0 + 0 = 0 but got %v", x)
	}

	x = BitAdd(234234, 345984560791)
	if x != 345984795025 {
		t.Fatalf("expected 234234 + 345984560791 = 345984795025 but got %v", x)
	}

}
