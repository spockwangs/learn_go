package util

import (
	"testing"
)

func TestMin(t *testing.T) {
	if ans := Min(1, 2); ans != 1 {
		t.Errorf("expected 1, but got %v", ans)
	}
}
