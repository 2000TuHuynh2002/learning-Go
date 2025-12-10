package main

import (
	"testing"
)

// Benchmark tests typically go in `_test.go` files and are named beginning with
// `Benchmark`. Any code that’s required for the benchmark to run but should not
// be measured goes before this loop.
func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		// The benchmark runner will automatically execute this loop body many times
		// to determine a reasonable estimate of the run-time of a single iteration.
		IntMin(1, 2)
	}
}
