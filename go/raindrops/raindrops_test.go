package raindrops

import "testing"

var tests = []struct {
	input    int
	expected string
}{
	{1, "1"},
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
	{6, "Pling"},
	{9, "Pling"},
	{10, "Plang"},
	{14, "Plong"},
	{15, "PlingPlang"},
	{21, "PlingPlong"},
	{25, "Plang"},
	{35, "PlangPlong"},
	{49, "Plong"},
	{52, "52"},
	{105, "PlingPlangPlong"},
	{12121, "12121"},
}

func TestConvert(t *testing.T) {
	for _, test := range tests {
		if actual := Convert(test.input); actual != test.expected {
			t.Errorf("Convert(%d) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Convert(test.input)
		}
	}
}

func TestPrime(t *testing.T) {
	assert(17, IsPrime(17), t)
	assert(2, IsPrime(2), t)
	assert(3, IsPrime(3), t)
	assert(8, !IsPrime(8), t)
}

func assert(n int, statement bool, t *testing.T) {
	if !statement {
		t.Errorf("assertion failed for %d", n)
	}
}
