package chapter01

import "testing"

func TestIsUnique(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"Zbcd", true},
		{"abcc", false},
		{" ", true},
		{"", true},
	}
	for _, c := range cases {
		actual := IsUnique(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %v, actual: %v\n", c.input, c.expected, actual)
		}
		actual = IsUniqueNoAllocation(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s. Expected: %v, actual: %v\n", c.input, c.expected, actual)
		}
	}
}
