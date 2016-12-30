package chapter01

// IsUnique solve the problem using an array of rune. Can be any acii chars
func IsUnique(str string) bool {
	if len(str) > 128 {
		return false
	}
	chars := make([]bool, 128)
	for _, runeVal := range str {
		if chars[runeVal] == true {
			return false
		}
		chars[runeVal] = true
	}
	return true
}

// IsUniqueNoAllocation solve the problem by using one byte. Can be only 'a' < x < 'z'
func IsUniqueNoAllocation(str string) bool {
	if len(str) > 26 {
		return false
	}
	checker := 0
	for _, runeVal := range str {
		val := uint(runeVal - 'a')
		if (checker & (1 << val)) > 0 {
			return false
		}
		checker |= (1 << val)
	}
	return true
}
