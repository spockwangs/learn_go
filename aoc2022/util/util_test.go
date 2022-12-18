package util

import (
	"testing"
)

func TestMin(t *testing.T) {
	if ans := MinInt(1, 2); ans != 1 {
		t.Errorf("expected 1, but got %v", ans)
	}
}

func TestConsumeStrUntil(t *testing.T) {
	s := "abc def"
	next := 0
	ss := ConsumeStrUntil(s, " ", &next)
	if ss != "abc" || next != 4 {
		t.Errorf("expected `abc`, but got `%v`", ss)
	}
	ss = ConsumeStrUntil(s, "  ", &next)
	if ss != "def" {
		t.Errorf("expected `def`, but got `%v`", ss)
	}
	if next != len(s) {
		t.Errorf("expected %v, but got %v", len(s), next)
	}
}
