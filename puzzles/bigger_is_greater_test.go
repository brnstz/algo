package puzzles

import "testing"

func TestBiggerIsGreater1(t *testing.T) {
	expected := "ba"
	x := BiggerIsGreater("ab")
	if x != expected {
		t.Fatalf("expected %v but got %v", expected, x)
	}
}

func TestBiggerIsGreater2(t *testing.T) {
	expected := "zedawdvyyfumwpupuinbdbfndyehircmylbaowuptgwm"
	x := BiggerIsGreater("zedawdvyyfumwpupuinbdbfndyehircmylbaowuptgmw")
	if x != expected {
		t.Fatalf("expected %v but got %v", expected, x)
	}
}

func TestBiggerIsGreater3(t *testing.T) {
	expected := "no answer"
	x := BiggerIsGreater("zyyxwwtrrnmlggfeb")
	if x != expected {
		t.Fatalf("expected %v but got %v", expected, x)
	}
}
