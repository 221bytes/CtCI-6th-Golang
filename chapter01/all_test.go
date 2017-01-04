package chapter01

import "testing"

func TestIsUnique(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"abcde", true},
		{"hello", false},
		{"apple", false},
		{"kite", true},
		{"padle", true},
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

func TestArePermutations(t *testing.T) {
	cases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"abcd", "abcd", true},
		{"abcd", "abdc", true},
		{"abcc", "ccbb", false},
		{"abcc", "abcc ", false},
		{" ", " ", true},
		{"", "", true},
	}
	for _, c := range cases {
		actual := ArePermuttationSort(c.input1, c.input2)
		if actual != c.expected {
			t.Fatalf("Inputs %s, %s. Expected: %v, actual: %v\n",
				c.input1, c.input2, c.expected, actual)
		}
		actual = ArePermuttationRuneCounter(c.input1, c.input2)
		if actual != c.expected {
			t.Fatalf("Inputs %s, %s. Expected: %v, actual: %v\n",
				c.input1, c.input2, c.expected, actual)
		}
	}
}

func isRunesSliceIsEqual(a, b []rune) bool {
	for index := 0; index < len(a) || index < len(b); index++ {
		if a[index] != b[index] {
			return false
		}
	}
	return true
}

func TestURLify(t *testing.T) {
	cases := []struct {
		input        []rune
		actualLength int
		expected     []rune
	}{
		{[]rune("Mr John Smith    "), 13, []rune("Mr%20John%20Smith")},
	}
	for _, c := range cases {
		actual := URLify(c.input, c.actualLength)
		if !isRunesSliceIsEqual(actual, c.expected) {
			t.Fatalf("Input %s. Expected: %v, actual: %v\n", c.input, string(c.expected), string(actual))
		}
	}
}
