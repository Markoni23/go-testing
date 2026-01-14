package even

import "testing"

func TestIsEven(t *testing.T) {
	if !isEven(2) {
		t.Error("2: expected - true, got - false")
	}
	if isEven(3) {
		t.Error("3: expected - false, got - true")
	}
	if !isEven(0) {
		t.Error("0: expected - true, got - false")
	}
	if !isEven(-4) {
		t.Error("-4: expected - true, got - false")
	}
	if isEven(-7) {
		t.Error("-7: expected - false, got - true")
	}
}
