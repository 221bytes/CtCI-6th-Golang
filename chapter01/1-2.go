package chapter01

import "sort"
import "strings"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// ArePermuttationSort sort the problem by sorting the string then by compare it
func ArePermuttationSort(str0, str1 string) bool {
	if len(str0) != len(str1) {
		return false
	}
	str0Sorted := sortString(str0)
	str1Sorted := sortString(str1)
	if strings.Compare(str0Sorted, str1Sorted) != 0 {
		return false
	}
	return true
}

// ArePermuttationRuneCounter use an array of int to count the runes
func ArePermuttationRuneCounter(str0, str1 string) bool {
	if len(str0) != len(str1) {
		return false
	}
	runesCounter := make([]int, 128)
	for _, runeVal := range str0 {
		runesCounter[runeVal]++
	}
	for _, runeVal := range str1 {
		runesCounter[runeVal]--
		if runesCounter[runeVal] < 0 {
			return false
		}
	}
	return true
}
