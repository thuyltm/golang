package main

import "testing"

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasix(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2_ = %d; want -2", ans)
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
