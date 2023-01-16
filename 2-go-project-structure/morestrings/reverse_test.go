package morestrings

import (
	"fmt"
	"testing"
)

func TestReverseRunes(t *testing.T) {

	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	fmt.Printf("struct=%v, type=%T\n", cases, cases)

	for i, c := range cases {
		got := ReverseRunes(c.in)
		fmt.Printf("(%v) c.in=[%v], want=[%v], got=[%v]\n", i, c.in, c.want, got)

		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

// Run the benchmark: go test -bench=BenchmarkReverseRunes
// Run multiple  benchmarks: go test -bench=BenchmarkReverseRunes -count INTEGER
// BenchmarkReverseRunes-2   	 4380300	       265.6 ns/op
// - total num of times is being executed, X ns/op It is the average time each iteration took to complete (nanos/operation)
// Adjusting the min time we want to be run: go test -bench=. -benchtime=10s
func BenchmarkReverseRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseRunes("my-text")
	}
}
