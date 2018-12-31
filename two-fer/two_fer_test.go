package twofer

import "testing"

// test cases to run on ShareWith function
var tests = []struct {
	name, expected string
}{
	{"", "One for you, one for me."},
	{"Alice", "One for Alice, one for me."},
	{"Bob", "One for Bob, one for me."},
}

// Function responsible for executing the test cases
func TestShareWith(t *testing.T) {
	for _, test := range tests {
		if observed := ShareWith(test.name); observed != test.expected {
			t.Fatalf("ShareWith(%s) = %v, want %v", test.name, observed, test.expected)
		}
	}
}

// Function responsible for running benchmarks
func BenchmarkShareWith(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, test := range tests {
			ShareWith(test.name)
		}

	}
}
